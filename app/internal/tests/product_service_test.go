package tests

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/models"
	"bzhspback.fr/breizhsport/internal/services"
)

// Mock repository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductByID(id uint) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductsByCategoryID(categoryID uint, filters []dto.Filter) ([]models.Product, error) {
	args := m.Called(categoryID, filters)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetRecommendedProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductsByNameOrDescription(nameOrDescription string, filters []dto.Filter) ([]models.Product, error) {
	args := m.Called(nameOrDescription, filters)
	return args.Get(0).([]models.Product), args.Error(1)
}

func TestGetProductByID(t *testing.T) {
	mockRepo := &MockProductRepository{}
	service := services.NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedProduct := models.Product{
			Model: gorm.Model{ID: 1},
			Title: "Test Product",
			Description: "Test Description",
			Price: 99.99,
			CategoryID: 1,
			BrandID: 1,
			Colors: "Red,Blue",
			Sizes: "M,L",
			PictureUrl: "test.jpg",
		}

		mockRepo.On("GetProductByID", uint(1)).Return(expectedProduct, nil)

		result, err := service.GetProductByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct.ID, result.ID)
		assert.Equal(t, expectedProduct.Title, result.Title)
		assert.Equal(t, expectedProduct.Description, result.Description)
		assert.Equal(t, expectedProduct.Price, result.Price)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("GetProductByID", uint(999)).Return(models.Product{}, errors.New("product not found"))

		_, err := service.GetProductByID(999)

		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())
	})
}

func TestGetProductsByCategoryID(t *testing.T) {
	mockRepo := &MockProductRepository{}
	service := services.NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedProducts := []models.Product{
			{
				Model: gorm.Model{ID: 1},
				Title: "Product 1",
				CategoryID: 1,
				Price: 99.99,
				Colors: "Red",
				Sizes: "M",
				PictureUrl: "product1.jpg",
			},
			{
				Model: gorm.Model{ID: 2},
				Title: "Product 2",
				CategoryID: 1,
				Price: 89.99,
				Colors: "Blue",
				Sizes: "L",
				PictureUrl: "product2.jpg",
			},
		}

		filterRequest := dto.FilterRequest{
			Filters: []dto.Filter{{Field: "price", Value: "100"}},
		}

		mockRepo.On("GetProductsByCategoryID", uint(1), filterRequest.Filters).Return(expectedProducts, nil)

		results, err := service.GetProductsByCategoryID(1, filterRequest)

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, expectedProducts[0].ID, results[0].ID)
		assert.Equal(t, expectedProducts[1].ID, results[1].ID)
	})
}

func TestGetRecommendedProducts(t *testing.T) {
	mockRepo := &MockProductRepository{}
	service := services.NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedProducts := []models.Product{
			{
				Model: gorm.Model{ID: 1},
				Title: "Recommended 1",
				Price: 99.99,
				Colors: "Red",
				Sizes: "M",
				PictureUrl: "rec1.jpg",
			},
			{
				Model: gorm.Model{ID: 2},
				Title: "Recommended 2",
				Price: 89.99,
				Colors: "Blue",
				Sizes: "L",
				PictureUrl: "rec2.jpg",
			},
		}

		mockRepo.On("GetRecommendedProducts").Return(expectedProducts, nil)

		results, err := service.GetRecommendedProducts()

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, expectedProducts[0].ID, results[0].ID)
		assert.Equal(t, expectedProducts[1].ID, results[1].ID)
	})
}

func TestGetProductsByNameOrDescription(t *testing.T) {
	mockRepo := &MockProductRepository{}
	service := services.NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedProducts := []models.Product{
			{
				Model: gorm.Model{ID: 1},
				Title: "Sport Shoes",
				Description: "Running shoes",
				Price: 99.99,
				Colors: "Black",
				Sizes: "42,43",
				PictureUrl: "shoes.jpg",
			},
			{
				Model: gorm.Model{ID: 2},
				Title: "Running Shorts",
				Description: "Sports shorts",
				Price: 29.99,
				Colors: "Black,Blue",
				Sizes: "M,L",
				PictureUrl: "shorts.jpg",
			},
		}

		filterRequest := dto.FilterRequest{
			Filters: []dto.Filter{{Field: "brand", Value: "Nike"}},
		}

		mockRepo.On("GetProductsByNameOrDescription", "sport", filterRequest.Filters).Return(expectedProducts, nil)

		results, err := service.GetProductsByNameOrDescription("sport", filterRequest)

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, expectedProducts[0].ID, results[0].ID)
		assert.Equal(t, expectedProducts[1].ID, results[1].ID)
	})
}
