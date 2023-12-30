package handler

import (
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		machines := api.Group("/machines")
		{
			machines.POST("/validate", h.validateMachines)
		}
	}

	return router
}
