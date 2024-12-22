package repository

import (
	"context"
	"errors"
	"fmt"

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
