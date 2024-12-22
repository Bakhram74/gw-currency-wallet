package v1

import (
	"github.com/Bakhram74/gw-currency-wallet/internal/service"

	"github.com/Bakhram74/gw-currency-wallet/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Router struct {
	service  *service.Service
	jwtMaker *jwt.JWTMaker
}

func NewRoute(service *service.Service, jwtMaker *jwt.JWTMaker) *Router {
	return &Router{
		jwtMaker: jwtMaker,
		service:  service,
	}
}

func (r *Router) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.POST("/register", r.register)
		v1.POST("/login", r.login)
	}

	v1.Use(authMiddleware(r.jwtMaker))
	{
		v1.GET("/balance", r.balance)
		v1.POST("/wallet/deposit", r.deposit)
	}
}
