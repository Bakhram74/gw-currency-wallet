package http

import (
	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	v1 "github.com/Bakhram74/gw-currency-wallet/internal/http/v1"
	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	"github.com/Bakhram74/gw-currency-wallet/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	config   config.Config
	service  *service.Service
	jwtMaker *jwt.JWTMaker
}

func NewHandler(config config.Config, service *service.Service, jwtMaker *jwt.JWTMaker) *Handler {

	return &Handler{
		service:  service,
		config:   config,
		jwtMaker: jwtMaker,
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
	v1 := v1.NewRoute(h.service, h.jwtMaker)
	{
		v1.Init(api)
	}
}
