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

func GetOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Oppening",
	})
}

func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Oppening",
	})
}

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Oppening",
	})
}

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE  Oppening",
	})
}