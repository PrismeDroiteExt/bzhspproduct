package dto

import (
	"time"

	"bzhspback.fr/breizhsport/internal/models"
)

type CategoryResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PictureUrl  string    `json:"picture_url"`
}

// ToResponse converts a Category model to a CategoryResponse
func ToResponse(category *models.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:          uint64(category.ID),
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
		PictureUrl:  category.PictureUrl,
	}
}

// ToResponseList converts a slice of Category models to a slice of CategoryResponse
func ToResponseList(categories []models.Category) []CategoryResponse {
	var responseList []CategoryResponse
	for _, category := range categories {
		categoryResponse := CategoryResponse{
			ID:          uint64(category.ID),
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
			PictureUrl:  category.PictureUrl,
		}
		responseList = append(responseList, categoryResponse)
	}
	return responseList
}
