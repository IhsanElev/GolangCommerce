package product

import "time"

type ProductListResponse struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductListResponseFromEntity(Product []Product) []ProductListResponse {
	var productList = []ProductListResponse{}
	for _, product := range Product {
		productList = append(productList, ProductListResponse{
			ID:    product.ID,
			SKU:   product.SKU,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}
	return productList
}

type ProductDetailResponse struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
