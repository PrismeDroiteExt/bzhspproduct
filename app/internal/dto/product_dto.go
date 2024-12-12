package dto

type ProductResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Discount    float64          `json:"discount"`
	Colors      string           `json:"colors"`
	Sizes       string           `json:"sizes"`
	Brand       BrandResponse    `json:"brand"`
	Category    CategoryResponse `json:"category"`
	PictureUrl  string           `json:"picture_url"`
}
