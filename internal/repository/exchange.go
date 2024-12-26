package repository

import (
	"context"
	"fmt"

	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
)

type ExchangeRepo struct {
	db DBTX
}

func NewExchangeRepo(db DBTX) *ExchangeRepo {
	return &ExchangeRepo{
		db: db,
	}
}

func (e *ExchangeRepo) ExchangeCurrency(ctx context.Context, userID, fromCurrency, toCurrency string, rate, amount float32) (entity.ExchangeRepoResponse, error) {

	var wallet Wallet
	walletQuery := `SELECT usd, rub, eur FROM wallet WHERE user_id = $1`
	err := e.db.QueryRow(ctx, walletQuery, userID).Scan(&wallet.Usd, &wallet.Rub, &wallet.Eur)
	if err != nil {
		return entity.ExchangeRepoResponse{}, fmt.Errorf("failed to fetch wallet: %w", err)
	}

	currentBalance := map[string]float32{"USD": wallet.Usd, "EUR": wallet.Eur, "RUB": wallet.Rub}[fromCurrency]
	if currentBalance < amount {
		return entity.ExchangeRepoResponse{}, ErrInsufficientBalance
	}

	exchangedAmount := amount * rate
	newBalance := map[string]float32{"USD": wallet.Usd, "EUR": wallet.Eur, "RUB": wallet.Rub}
	newBalance[fromCurrency] -= amount
	newBalance[toCurrency] += exchangedAmount

	updateQuery := `UPDATE wallet SET usd = $1, rub = $2, eur = $3 WHERE user_id = $4`
	_, err = e.db.Exec(ctx, updateQuery, newBalance["USD"], newBalance["RUB"], newBalance["EUR"], userID)

	if err != nil {
		return entity.ExchangeRepoResponse{}, fmt.Errorf("failed to update wallet: %w", err)
	}

	return entity.ExchangeRepoResponse{
		ExchangedAmount: exchangedAmount,
		FromBalance:     newBalance[fromCurrency],
		ToBalance:       newBalance[toCurrency],
	}, nil

}
