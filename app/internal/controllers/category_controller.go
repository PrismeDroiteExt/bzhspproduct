package controllers

import (
	"net/http"
	"strconv"

	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

// GetAllCategories handles the request to get all categories
func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := cc.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch categories",
		})
		return
	}

	response := dto.ToResponseList(categories)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}

// GetCategoryByID handles the request to get a specific category by ID
func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid category ID",
		})
		return
	}

	category, err := cc.service.GetCategoryByID(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Category not found",
		})
		return
	}

	response := dto.ToResponse(category)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}
