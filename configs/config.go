package configs

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	API                  string        `mapstructure:"API"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
	PostgresURI          string        `mapstructure:"POSTGRES_URI"`
	RedisHost            string        `mapstructure:"REDIS_HOST"`
	RedisPort            string        `mapstructure:"REDIS_PORT"`
	JSONAPIHost          string        `mapstructure:"JSON_API_HOST"`
	JSONAPIPort          string        `mapstructure:"JSON_API_PORT"`
	GRPCAPIHost          string        `mapstructure:"GRPC_API_HOST"`
	GRPCAPIPort          string        `mapstructure:"GRPC_API_PORT"`
	EmailSenderName      string        `mapstructure:"EMAIL_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_PASSWORD"`
	SymmetricTokenKey    string        `mapstructure:"SYMMETRIC_TOKEN_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
