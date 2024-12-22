package service

import (
	"context"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/jwt"
)

type Auth interface {
	Register(ctx context.Context, username, password, email string) error
	Login(ctx context.Context, username, password string) (string, error)
}

type Balance interface {
	GetBalance(ctx context.Context, userID string) (repository.Wallet, error)
	DepositBalance(ctx context.Context,userID string,param entity.Transaction)(repository.Wallet, error)
	WithdrawBalance(ctx context.Context,userID string,param entity.Transaction)(repository.Wallet, error)
}

type Service struct {
	Auth
	Balance
}

func NewService(repo *repository.Repository, JwtMaker *jwt.JWTMaker, cfg config.Config) *Service {

	return &Service{
		Auth:    NewAuthService(repo, JwtMaker, cfg),
		Balance: NewBalanceService(repo, cfg),
	}

}
