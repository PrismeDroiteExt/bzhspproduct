package api

import (
	"bzhspback.fr/breizhsport/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup, controller *controllers.CategoryController) {
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("", controller.GetAllCategories)
		categoryRoutes.GET("/:id", controller.GetCategoryByID)
	}
}
