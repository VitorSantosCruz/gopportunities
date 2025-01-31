package handler

import (
	"net/http"

	"github.com/VitorSantosCruz/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error get-openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error get-openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "get-openings", openings)
}
