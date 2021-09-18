package http

import (
	"github.com/gin-gonic/gin"

	"github.com/mehrdod/grpcsimple/internal/service"
	"github.com/mehrdod/grpcsimple/pkg/logger"
)

type Handler struct {
	services *service.Services
	logger   logger.Logger
}

func NewHandler(services *service.Services, logger logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
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
	router.GET("/fibonacci", h.fibonacciGet)
}
