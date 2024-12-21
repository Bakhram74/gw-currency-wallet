package config

import (
	"log"
	"template/pkg/env"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	Storage StorageConfig
	JWT     TokenConfig
	Http    Http
}

type Http struct {
	Port string
}

type StorageConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
	PostgresSslMode  bool
}

type TokenConfig struct {
	TokenSecretKey      string
	AccessTokenDuration time.Duration
}

func NewConfig() Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading config.env file")
	}
	storage := StorageConfig{
		PostgresHost:     env.GetEnv("HOST_DB", "localhost"),
		PostgresPort:     env.GetEnv("PORT_DB", "5432"),
		PostgresDatabase: env.GetEnv("DATABASE", "wallet"),
		PostgresUsername: env.GetEnv("USERNAME_DB", "postgres"),
		PostgresPassword: env.GetEnv("PASSWORD_DB", "secret"),
		PostgresSslMode:  env.GetEnvAsBool("SSL_MODE", false),
	}

	config := Config{
		Env:     env.GetEnv("ENVIRONMENT", "local"),
		Storage: storage,
	}
	return config
}
