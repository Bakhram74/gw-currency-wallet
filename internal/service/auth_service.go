package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/pkg/logs"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils/jwt"
)

type AuthService struct {
	repo     *repository.Repository
	JwtMaker *jwt.JWTMaker
	cfg      config.Config
}

func NewAuthService(repo *repository.Repository, JwtMaker *jwt.JWTMaker, cfg config.Config) *AuthService {
	return &AuthService{
		repo:     repo,
		JwtMaker: JwtMaker,
		cfg:      cfg,
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
	user, err := a.repo.UserQueries.CreateUser(ctx, username, hashedPassword, email)
	if err != nil {
		log.Error("failed to create user", logs.Err(err))
		return err
	}

	if err := a.repo.WalletQueries.CreateWallet(ctx, user.ID); err != nil {
		log.Error("Failed to create wallet for user", logs.Err(err))
	}
	return nil
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	const op = "Auth.Login"

	log := slog.With(
		slog.String("op", op),
		slog.String("username", username),
	)
	log.Info("attempting to get user")

	user, err := a.repo.UserQueries.GetUser(ctx, username)
	if err != nil {
		log.Error("failed to get user", logs.Err(err))
		return "", err
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		log.Error("failed to get user", logs.Err(err))
		return "", repository.ErrUserNotFound
	}

	accessToken, _, err := a.JwtMaker.CreateToken(
		user,
		a.cfg.JWT.AccessTokenDuration,
	)
	if err != nil {
		log.Error("failed to create jwt token", logs.Err(err))
		return "", err
	}

	return accessToken, nil
}
