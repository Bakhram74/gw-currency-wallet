package jwt

import (
	"errors"

	"time"

	"github.com/Bakhram74/gw-currency-wallet/internal/repository"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        string    `json:"id"`
	Name      string    `json:"username,omitempty"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific user and duration
func NewPayload(user repository.User, duration time.Duration) *Payload {

	payload := &Payload{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
