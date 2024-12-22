package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/http"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	"github.com/Bakhram74/gw-currency-wallet/pkg/client/postgres"
	httpserver "github.com/Bakhram74/gw-currency-wallet/pkg/httpServer"
)

func Run(cfg config.Config) {
	dbUrl := url(cfg)

	pg, err := postgres.New(dbUrl)
	if err != nil {
		panic(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	err = RunMigration(dbUrl)
	if err != nil {
		panic(fmt.Sprintf("Migration error: %s", err.Error()))
	}

	repo := repository.New(pg.Pool)

	service := service.NewService(repo)

	handler := http.NewHandler(cfg, service).Init()

	slog.Debug("server starting", slog.String("port", cfg.HttpPort))
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HttpPort))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		slog.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	case err := <-httpServer.Notify():
		slog.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err).Error())
	}

	err = httpServer.Shutdown()
	if err != nil {
		slog.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}

func url(cfg config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Storage.PostgresUsername,
		cfg.Storage.PostgresPassword,
		cfg.Storage.PostgresHost,
		cfg.Storage.PostgresPort,
		cfg.Storage.PostgresDatabase,
		cfg.Storage.PostgresSslMode)
}
