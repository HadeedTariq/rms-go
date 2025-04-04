package routes

import (
	"rms-platform/controllers"
	"rms-platform/middleware"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(r *gin.RouterGroup) {
	menuGroup := r.Group("/menu")
	{
		menuGroup.Use(middleware.AuthMiddleware())
		menuGroup.Use(middleware.IsManagerOrAdmin())
		menuGroup.POST("/create", controllers.CreateMenu)
	}
}
