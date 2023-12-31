package handler

import (
	_ "github.com/bsydorenko7/machines-management-validator-backend/docs"
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		machines := api.Group("/machines")
		{
			machines.POST("/validate", h.validateMachines)
		}
	}

	return router
}
