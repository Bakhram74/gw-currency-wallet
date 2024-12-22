package service

import (
	"context"
	"log/slog"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
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
	const op = "Balance.GetBalance"

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

func (b *BalanceService) DepositBalance(ctx context.Context, userID string, param entity.Transaction) (repository.Wallet, error) {
	const op = "Balance.DepositBalance"

	log := slog.With(
		slog.String("op", op),
		slog.String("userID", userID),
		slog.Float64("amount", float64(param.Amount)),
		slog.String("currency", string(param.Currency)),
	)

	log.Info("attempting to deposit balance")

	wallet, err := b.repo.WalletQueries.DepositWallet(ctx, userID, string(param.Currency), param.Amount)
	if err != nil {
		log.Error("failed to deposit balance", logs.Err(err))
		return repository.Wallet{}, err
	}
	return wallet, nil
}

func (b *BalanceService) WithdrawBalance(ctx context.Context, userID string, param entity.Transaction) (repository.Wallet, error) {
	const op = "Balance.WithdrawBalance"

	log := slog.With(
		slog.String("op", op),
		slog.String("userID", userID),
		slog.Float64("amount", float64(param.Amount)),
		slog.String("currency", string(param.Currency)),
	)

	log.Info("attempting to withdraw balance")
	wallet, err := b.repo.WalletQueries.WithdrawWallet(ctx, userID, string(param.Currency), param.Amount)
	if err != nil {
		log.Error("failed to withdraw balance", logs.Err(err))
		return repository.Wallet{}, err
	}
	return wallet, nil
}
