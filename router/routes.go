package router

import (
	"github.com/fabiopasilva1/gooportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		// ?Show Openings
		v1.GET("/openings/:id", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}
}
