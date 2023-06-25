package grpc_api

import (
	"fmt"

	"bookkeeper/configs"
	"bookkeeper/internal/db"
	"bookkeeper/internal/protogen"
	"bookkeeper/internal/tokens"
	"bookkeeper/internal/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	protogen.UnimplementedBookkeeperServer
	config          configs.Config
	store           db.Store
	tokenMaker      tokens.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config configs.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := tokens.NewPasetoMaker(config.SymmetricTokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
