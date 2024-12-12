package services

import (
	"bzhspback.fr/breizhsport/internal/models"
	"bzhspback.fr/breizhsport/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetCategoryByID(id uint64) (*models.Category, error) {
	return s.repo.FindByID(id)
}
