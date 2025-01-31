package handler

import (
	"net/http"

	"github.com/VitorSantosCruz/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Basepath api/v1
// @Summary get openings
// @Description get all the job openings
// @tags openings
// @Accept json
// @Produce json
// @Success 200 {object} GetOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error get-openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error get-openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "get-openings", openings)
}
