package handler

import (
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/dataaccess/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Validate
// @Tags machines
// @Description validate list of machines and return only outlier machines
// @Accept  json
// @Produce  json
// @Param input body []model.Machine true "list of machines"
// @Success 200 {object} []model.Machine
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/machines/validate [post]
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
