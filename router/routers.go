package router

import (
	"github.com/VitorSantosCruz/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")

	{
		v1.GET("/openings", handler.GetOpeningsHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	}
}
