package handler

import (
	"fmt"
	"net/http"

	"github.com/VitorSantosCruz/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id:  %s not found", id))
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating  opening in database")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-opening", opening)
}
