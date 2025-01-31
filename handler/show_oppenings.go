package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOppeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Oppenings",
	})
}