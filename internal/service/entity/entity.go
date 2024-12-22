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
