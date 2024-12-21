package v1

import (
	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	service *service.Service
}

func NewRoute(service *service.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.POST("/register", r.register)
		v1.POST("/login", r.login)
	}
}
