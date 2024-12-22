package entity

type RegisterReq struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginReq struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Balance struct {
	USD float32 `json:"USD"`
	RUB float32 `json:"RUB"`
	EUR float32 `json:"EUR"`
}
