package service

import (
	"context"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils/jwt"
)

type Auth interface {
	Register(ctx context.Context, username, password, email string) error
	Login(ctx context.Context, username, password string) (string, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository, JwtMaker *jwt.JWTMaker, cfg config.Config) *Service {

	return &Service{
		Auth: NewAuthService(repo, JwtMaker, cfg),
	}

}
