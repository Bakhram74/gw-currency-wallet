package app

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 5
	defaultTimeout  = time.Second
)

func RunMigration(url string) error {

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", url)
		if err == nil {
			break
		}

		slog.Debug("Migrate: postgres is trying to connect,", slog.Int("attempts left", attempts))
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		slog.Info("Migrate: no change")
		return nil
	}

	if err != nil {
		return tryCloseMigrateWithError(m, err)
	}

	slog.Info("Migrate: up success")

	return tryCloseMigrateWithError(m, nil)
}

func tryCloseMigrateWithError(m *migrate.Migrate, err error) error {
	var resultErr error
	sourceErr, databaseErr := m.Close()
	if sourceErr != nil {
		resultErr = fmt.Errorf("failed to close source, err: %w", sourceErr)
	}
	if databaseErr != nil {
		resultErr = errors.Join(resultErr, fmt.Errorf("failed to close database, err: %w", databaseErr))
	}
	return errors.Join(err, resultErr)
}
