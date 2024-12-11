package seeders

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

func SeedBrands(db *gorm.DB) error {
	brands := []models.Brand{
		{
			Name: "Nike",
		},
		{
			Name: "Adidas",
		},
		{
			Name: "Puma",
		},
	}

	for _, brand := range brands {
		result := db.FirstOrCreate(&brand, models.Brand{
			Name: brand.Name,
		})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
