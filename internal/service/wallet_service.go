package service

import (
	"context"
	"log/slog"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/pkg/logs"
)

type BalanceService struct {
	repo *repository.Repository
	cfg  config.Config //TODO delete?
}

func NewBalanceService(repo *repository.Repository, cfg config.Config) *BalanceService {
	return &BalanceService{
		repo: repo,
		cfg:  cfg,
	}
}

func (b *BalanceService) GetBalance(ctx context.Context, userID string) (repository.Wallet, error) {
	const op = "Auth.Register"

	log := slog.With(
		slog.String("op", op),
		slog.String("userID", userID),
	)
	log.Info("attempting to get balance")

	wallet, err := b.repo.WalletQueries.GetWallet(ctx, userID)
	if err != nil {
		log.Error("failed to get user balance", logs.Err(err))
		return repository.Wallet{}, err
	}

	return wallet, nil
}
