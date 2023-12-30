package handler

import (
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/dataaccess/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) validateMachines(c *gin.Context) {
	var input []model.Machine

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	
	machines, err := h.services.ValidateMachines(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, machines)
}
