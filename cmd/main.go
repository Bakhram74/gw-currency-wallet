package main

import (
	"log/slog"
	"os"

	"github.com/Bakhram74/gw-currency-wallet/internal/app"
	"github.com/Bakhram74/gw-currency-wallet/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// @title      Wallet-exchanger
// @version 0.0.1
// @description     API docs for Wallet-exchanger
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	cfg := config.NewConfig()
	log := setupLogger(cfg.Env)
	slog.SetDefault(log)

	app.Run(cfg)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
