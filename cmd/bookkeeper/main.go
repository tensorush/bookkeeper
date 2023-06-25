package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"bookkeeper/configs"
	_ "bookkeeper/dbdocs/statik"
	"bookkeeper/internal/db"
	"bookkeeper/internal/grpc_api"
	"bookkeeper/internal/json_api"
	"bookkeeper/internal/notifications"
	"bookkeeper/internal/protogen"
	"bookkeeper/internal/worker"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := configs.LoadConfig("./configs/")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration("file://internal/db/migrations/", config.PostgresURI)

	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	if config.API == "json" {
		runJSONServer(config, store)
	} else {
		taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
		go runTaskProcessor(config, redisOpt, store)
		go runGatewayServer(config, store, taskDistributor)
		runGrpcServer(config, store, taskDistributor)
	}
}

func runDBMigration(migrationPath string, PostgresURI string) {
	migrations, err := migrate.New(migrationPath, PostgresURI)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migrations.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runTaskProcessor(config configs.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := notifications.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGrpcServer(config configs.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := grpc_api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(grpc_api.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	protogen.RegisterBookkeeperServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCAPIAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(config configs.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := grpc_api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = protogen.RegisterBookkeeperHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.JSONAPIAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := grpc_api.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}

func runJSONServer(config configs.Config, store db.Store) {
	server, err := json_api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.JSONAPIAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
