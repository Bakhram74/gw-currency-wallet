package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserExists   = errors.New("username or email already exists")
	ErrEmailFormat  = errors.New("invalid email value")
	ErrUserNotFound = errors.New("invalid username or password")
)

const (
	UniqueViolation = "23505"
	CheckViolation  = "23514"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type UserQueries interface {
	CreateUser(ctx context.Context, username, password, email string) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
}

type Repository struct {
	connPool *pgxpool.Pool
	UserQueries
}

func New(connPool *pgxpool.Pool) *Repository {
	return &Repository{
		connPool:    connPool,
		UserQueries: NewUserRepo(connPool),
	}
}
