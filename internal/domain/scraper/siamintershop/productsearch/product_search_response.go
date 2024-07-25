package productsearch

import (
	"manga-crawler/internal/domain/catalogupdate/product"

	"github.com/google/uuid"
)

type ProductSearchResponse struct {
	Total    string                         `json:"total"`
	Limit    string                         `json:"limit"`
	Offset   int                            `json:"offset"`
	Products []ProductSearchResponseProduct `json:"products"`
}

type ProductSearchResponseProduct struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`

	ProductMinPrice    float64 `json:"product_min_price"`
	ProductMaxPrice    float64 `json:"product_max_price"`
	ProductFullPrice   float64 `json:"product_full_price"`
	ProductSalePercent float64 `json:"product_sale_percent"`
}

func (p ProductSearchResponseProduct) ToProduct(publisherID uuid.UUID) (*product.Product, error) {
	return product.NewProduct(product.NewProductArgs{
		Name:        p.ProductName,
		ExternalID:  p.ProductId,
		Price:       p.ProductMaxPrice,
		Discounted:  p.isDiscounted(),
		PublisherID: publisherID,
	})
}

func (p ProductSearchResponseProduct) isDiscounted() bool {
	if p.ProductSalePercent > 0 {
		return true
	}

	return p.ProductMaxPrice != p.ProductFullPrice
}
