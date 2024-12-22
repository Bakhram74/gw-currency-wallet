package config

import (
	"log"
	"time"

	"github.com/Bakhram74/gw-currency-wallet/pkg/env"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Storage  StorageConfig
	JWT      TokenConfig
	HttpPort string
}

type StorageConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
	PostgresSslMode  string
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
		PostgresSslMode:  env.GetEnv("SSL_MODE", "disable"),
	}
	token := TokenConfig{
		TokenSecretKey:      env.GetEnv("TOKEN_SECRET_KEY", "cdwasfr43q12deasw90fj32lf8snre13"),
		AccessTokenDuration: time.Hour * 100,
	}
	config := Config{
		JWT:      token,
		HttpPort: env.GetEnv("HTTP_PORT", "9090"),
		Env:      env.GetEnv("ENVIRONMENT", "local"),
		Storage:  storage,
	}
	return config
}
