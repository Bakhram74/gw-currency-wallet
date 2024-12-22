package repository

import (
	"context"
	"fmt"
)

type WalletRepo struct {
	db DBTX
}

func NewWalletRepo(db DBTX) *WalletRepo {
	return &WalletRepo{
		db: db,
	}
}

func (w *WalletRepo) CreateWallet(ctx context.Context, userID string) error {
	query := `INSERT INTO "wallet" (user_id) values ($1);`

	_, err := w.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("could not create wallet: %w", err)
	}
	return nil
}

func (w *WalletRepo) GetWallet(ctx context.Context, userID string) (Wallet, error) {
	fmt.Println("ID" + userID)
	query := `SELECT user_id, usd,rub, eur FROM wallet WHERE user_id = $1`

	row := w.db.QueryRow(ctx, query, userID)
	var i Wallet
	err := row.Scan(
		&i.UserID,
		&i.Usd,
		&i.Rub,
		&i.Eur,
	)
	if err != nil {
		return i, err
	}

	return i, nil
}
