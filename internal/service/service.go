package service

import (
	"context"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
)

type Auth interface {
	Register(ctx context.Context, username, password, email string) error
	// Login(username, password string) (string, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		Auth: NewAuthService(repo),
	}

}
