package v1

import (
	"errors"

	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	"github.com/Bakhram74/proto-exchange/pb"

	"github.com/Bakhram74/gw-currency-wallet/pkg/jwt"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidCurrencyAmount = errors.New("invalid amount or currency")
	ErrInsufficientAmount    = errors.New("insufficient funds or invalid amount")
)

type Router struct {
	service    *service.Service
	jwtMaker   *jwt.JWTMaker
	grpcClient pb.ExchangeServiceClient
}

func NewRoute(service *service.Service, jwtMaker *jwt.JWTMaker, grpcClient pb.ExchangeServiceClient) *Router {
	return &Router{
		jwtMaker:   jwtMaker,
		service:    service,
		grpcClient: grpcClient,
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
		v1.POST("/wallet/withdraw", r.withdraw)
	}
	{
		v1.GET("/exchange/rates", r.rates)
		v1.POST("/exchange", r.exchange)
	}

}
