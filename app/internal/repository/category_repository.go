package repository

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// FindAll retrieves all categories from the database
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}

// FindByID retrieves a category by its ID from the database
func (r *CategoryRepository) FindByID(id uint64) (*models.Category, error) {
	var category models.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
