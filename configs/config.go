package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/spf13/cast"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBUser          string
	DBName          string
	DBPassword      string
	AuthServicePort string
	AuthRouterPort  string
	ServiceName     string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file not found")
	}

	cfg := Config{}

	cfg.DBHost = cast.ToString(getOrReturnDefault("DB_HOST", "localhost"))
	cfg.DBPort = cast.ToString(getOrReturnDefault("DB_PORT", "5432"))
	cfg.DBUser = cast.ToString(getOrReturnDefault("DB_USER", "postgres"))
	cfg.DBName = cast.ToString(getOrReturnDefault("DB_NAME", "postgres"))
	cfg.DBPassword = cast.ToString(getOrReturnDefault("DB_PASSWORD", "123456789"))
	cfg.AuthServicePort = cast.ToString(getOrReturnDefault("AUTH_SERVICE_PORT", ":1111"))
	cfg.AuthRouterPort = cast.ToString(getOrReturnDefault("AUTH_ROUTER_PORT", ":2222"))
	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "nothing"))

	return &cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
