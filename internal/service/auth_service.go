package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/pkg/logs"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) Register(ctx context.Context, username, password, email string) error {
	const op = "Auth.Register"

	log := slog.With(
		slog.String("op", op),
		slog.String("email", email),
		slog.String("username", username),
	)
	log.Info("attempting to create user")

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Error("failed to generate password hash", logs.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = a.repo.UserQueries.CreateUser(ctx, username, hashedPassword, email)
	if err != nil {
		log.Error("failed to create user", logs.Err(err))
		return err
	}

	return nil
}

// func (a *AuthService) Login(username, password string) (string, error) {

// }
