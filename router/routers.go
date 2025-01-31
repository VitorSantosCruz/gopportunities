package router

import (
	"github.com/VitorSantosCruz/gopportunities/handler"
	"github.com/gin-gonic/gin"
	docs "github.com/VitorSantosCruz/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	{
		v1.GET("/openings", handler.GetOpeningsHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
