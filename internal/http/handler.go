package http

import (
	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	v1 "github.com/Bakhram74/gw-currency-wallet/internal/http/v1"
	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	config  config.Config
	service *service.Service
}

func NewHandler(config config.Config, service *service.Service) *Handler {

	return &Handler{
		service: service,
		config:  config,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	api := router.Group("/api")
	v1 := v1.NewRoute(h.service)
	{
		v1.Init(api)
	}
}
