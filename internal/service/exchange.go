package service

import (
	"context"
	"log/slog"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/Bakhram74/gw-currency-wallet/pkg/logs"
)

type ExchangeService struct {
	repo *repository.Repository
}

func NewExchangeService(repo *repository.Repository) *ExchangeService {
	return &ExchangeService{
		repo: repo,
	}
}

func (e *ExchangeService) ExchangeCurrency(ctx context.Context, userID, fromCurrency, toCurrency string, rate, amount float32) (entity.ExchangeRepoResponse, error) {
	const op = "Exchange.ExchangeCurrency"

	log := slog.With(
		slog.String("op", op),
		slog.String("userID", userID),
		slog.String("fromCurrency", fromCurrency),
		slog.String("toCurrency", toCurrency),
		slog.Float64("amount", float64(amount)),
	)
	log.Info("attempting exchange currency")

	exchanged, err := e.repo.ExchangeQueries.ExchangeCurrency(ctx, userID, fromCurrency, toCurrency, rate, amount)
	if err != nil {
		log.Error("failed to exchange currency", logs.Err(err))
		return entity.ExchangeRepoResponse{}, err
	}
	return exchanged, nil
}
