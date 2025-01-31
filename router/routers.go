package router

import (
	"github.com/VitorSantosCruz/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/oppenings", handler.GetOppeningsHandler)
		v1.GET("/oppenings/:id", handler.GetOppeningHandler)
		v1.POST("/oppenings", handler.CreateOppeningHandler)
		v1.PUT("/oppenings", handler.UpdateOppeningHandler)
		v1.DELETE("/oppenings/:id", handler.DeleteOppeningHandler)
	}
}
