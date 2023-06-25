package grpc_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	"bookkeeper/configs"
	"bookkeeper/internal/db"
	"bookkeeper/internal/tokens"
	"bookkeeper/internal/utils"
	"bookkeeper/internal/worker"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := configs.Config{
		SymmetricTokenKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker tokens.Maker, username string, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}

	return metadata.NewIncomingContext(context.Background(), md)
}
