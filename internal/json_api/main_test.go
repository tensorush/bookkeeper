package json_api

import (
	"os"
	"testing"
	"time"

	"bookkeeper/configs"
	"bookkeeper/internal/db"
	"bookkeeper/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := configs.Config{
		SymmetricTokenKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
