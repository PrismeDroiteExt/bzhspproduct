package seeders

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{
			Name:        "Sneakers",
			Description: "Athletic and casual footwear for all occasions",
			PictureUrl:  "/images/categories/sneakers.jpg",
		},
		{
			Name:        "T-Shirts",
			Description: "Comfortable and stylish t-shirts for everyday wear",
			PictureUrl:  "/images/categories/tshirts.jpg",
		},
		{
			Name:        "Shorts",
			Description: "Athletic and casual shorts for sports and leisure",
			PictureUrl:  "/images/categories/shorts.jpg",
		},
	}

	for _, category := range categories {
		result := db.FirstOrCreate(&category, models.Category{
			Name: category.Name,
		})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
