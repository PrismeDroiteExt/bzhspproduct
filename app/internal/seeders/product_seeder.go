package seeders

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	products := []models.Product{
		{
			Title:       "Nike Air Max 270",
			Description: "The Nike Air Max 270 delivers visible cushioning under every step. Updated for modern comfort, it features Nike's biggest heel Air unit yet.",
			CategoryID:  1, // Sneakers
			BrandID:     1, // Nike
			Price:       149.99,
			Discount:    nil,
			Colors:      "Black,White,Red",
			Sizes:       "40,41,42,43,44,45",
			PictureUrl:  "/images/products/nike.jpg",
		},
		{
			Title:       "Adidas Ultraboost 21",
			Description: "Experience epic energy with the adidas Ultraboost 21. The shoe features Boost cushioning for incredible energy return.",
			CategoryID:  1, // Sneakers
			BrandID:     2, // Adidas
			Price:       179.99,
			Discount:    newFloat64(159.99),
			Colors:      "White,Black,Blue",
			Sizes:       "39,40,41,42,43,44",
			PictureUrl:  "/images/products/adidas.jpg",
		},
		{
			Title:       "Puma RS-X",
			Description: "The Puma RS-X reinvents the classic Running System (RS) line with modern materials and bold color combinations.",
			CategoryID:  1, // Sneakers
			BrandID:     3, // Puma
			Price:       129.99,
			Discount:    newFloat64(119.99),
			Colors:      "White,Black,Yellow",
			Sizes:       "40,41,42,43,44",
			PictureUrl:  "/images/products/puma.jpg",
		},
		{
			Title:       "Nike Dri-FIT T-Shirt",
			Description: "Stay cool and dry during your workout with this Nike Dri-FIT technology t-shirt.",
			CategoryID:  2, // T-Shirts
			BrandID:     1, // Nike
			Price:       29.99,
			Discount:    nil,
			Colors:      "Black,White,Grey",
			Sizes:       "S,M,L,XL",
			PictureUrl:  "/images/products/nike-tshirt.jpg",
		},
		{
			Title:       "Adidas Training Shorts",
			Description: "Comfortable training shorts with AEROREADY technology to keep you dry during intense workouts.",
			CategoryID:  3, // Shorts
			BrandID:     2, // Adidas
			Price:       34.99,
			Discount:    newFloat64(10.0),
			Colors:      "Black,Navy,Grey",
			Sizes:       "S,M,L,XL",
			PictureUrl:  "/images/products/adidas-shorts.jpg",
		},
	}

	for _, product := range products {
		result := db.FirstOrCreate(&product, models.Product{
			Title:   product.Title,
			BrandID: product.BrandID,
		})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

// Helper function to create a pointer to a float64
func newFloat64(v float64) *float64 {
	return &v
}
