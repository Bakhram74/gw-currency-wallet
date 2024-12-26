package entity

type ExchangeReq struct {
	From   string  `json:"from_currency" binding:"required"`
	To     string  `json:"to_currency" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
}

type ExchangeResponse struct {
	Message    string             `json:"message"`
	Amount     float32            `json:"exchanged_amount"`
	NewBalance map[string]float32 `json:"new_balance"`
}

type ExchangeRepoResponse struct {
	ExchangedAmount float32 `json:"exchanged_amount"`
	FromBalance     float32 `json:"from_balance"`
	ToBalance       float32 `json:"to_balance"`
}
