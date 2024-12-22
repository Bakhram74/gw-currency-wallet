package entity

type Currency string

const (
	USD Currency = "USD"
	RUB Currency = "RUB"
	EUR Currency = "EUR"
)

type Transaction struct {
	Amount   float32  `json:"amount"`
	Currency Currency `json:"currency"`
}

func IsValidCurrency(c string) bool {
	switch Currency(c) {
	case USD, RUB, EUR:
		return true
	}
	return false
}

type Balance struct {
	USD float32 `json:"USD"`
	RUB float32 `json:"RUB"`
	EUR float32 `json:"EUR"`
}

type DepositResponse struct {
	Message    string  `json:"message"`
	NewBalance Balance `json:"new_balance"`
}
