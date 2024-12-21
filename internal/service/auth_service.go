package service

import (
	"context"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
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

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {

	}
	_, err = a.repo.UserQueries.CreateUser(ctx, username, hashedPassword, email)
	if err != nil {
		return err
	}

	return nil
}

// func (a *AuthService) Login(username, password string) (string, error) {

// }
