package repository

import (
	"context"
	"fmt"
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

	query := "INSERT INTO user (username,password,email) values ($1, $2, $3)"

	_, err := u.db.Exec(context.Background(), query, username, password, email)
	if err != nil {
		fmt.Println(err.Error()) //TODO check
	}
	return User{}, nil
}
