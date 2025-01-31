package handler

import (
	"fmt"
	"net/http"

	"github.com/VitorSantosCruz/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Basepath api/v1
// @Summary get opening
// @Description get a new job opening
// @tags openings
// @Accept json
// @Produce json
// @Param id path int true "Opening identification"
// @Success 200 {object} GetOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings/{id} [get]
func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		logger.Errorf("Error get opening: id is empty")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error get opening: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id:  %s not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "get-opening", opening)
}
