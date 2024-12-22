package repository

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Wallet struct {
	UserID string  `json:"user_id"`
	Usd    float32 `json:"usd"`
	Rub    float32 `json:"rub"`
	Eur    float32 `json:"eur"`
}
