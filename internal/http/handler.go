package http

import (
	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	v1 "github.com/Bakhram74/gw-currency-wallet/internal/http/v1"
	"github.com/Bakhram74/gw-currency-wallet/internal/service"
	_ "github.com/Bakhram74/gw-currency-wallet/docs"
	"github.com/Bakhram74/gw-currency-wallet/pkg/jwt"
	"github.com/Bakhram74/proto-exchange/pb"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	config     config.Config
	service    *service.Service
	jwtMaker   *jwt.JWTMaker
	grpcClient pb.ExchangeServiceClient
}

func NewHandler(config config.Config, service *service.Service, jwtMaker *jwt.JWTMaker, grpcClient pb.ExchangeServiceClient) *Handler {

	return &Handler{
		service:    service,
		config:     config,
		jwtMaker:   jwtMaker,
		grpcClient: grpcClient,
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
	v1 := v1.NewRoute(h.service, h.jwtMaker, h.grpcClient)
	{
		v1.Init(api)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
