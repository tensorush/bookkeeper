package db

import (
	"bookkeeper/configs"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := configs.LoadConfig("../../configs/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
