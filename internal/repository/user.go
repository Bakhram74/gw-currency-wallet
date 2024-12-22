package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepo struct {
	db DBTX
}

func NewUserRepo(db DBTX) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, username, password, email string) (User, error) {

	query := `INSERT INTO "user" (username,password,email)
	 values ($1, $2, $3) RETURNING id,username,email,created_at`

	row := u.db.QueryRow(context.Background(), query, username, password, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
	)
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case CheckViolation:
				return i, ErrEmailFormat

			case UniqueViolation:
				return i, ErrUserExists
			}
		}

		return i, fmt.Errorf("failed to create user: %w", err)
	}
	return i, nil

}

func (u *UserRepo) GetUser(ctx context.Context, username string) (User, error) {
	query := `SELECT id, username,password, email,created_at FROM "user"
      WHERE username = $1`

	row := u.db.QueryRow(ctx, query, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return i, ErrUserNotFound
		}
		return i, fmt.Errorf("failed to get user: %w", err)
	}
	return i, nil
}
