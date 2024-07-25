package productdetail

import (
	"manga-crawler/internal/domain/catalogupdate/product"

	"github.com/google/uuid"
)

type ProductDetail struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductTitle string `json:"product_title"`

	ProductMinPrice  float64 `json:"product_min_price"`
	ProductMaxPrice  float64 `json:"product_max_price"`
	ProductFullPrice float64 `json:"product_full_price"`

	ProductSalePercent float64 `json:"product_sale_percent"`

	Flag       string    `json:"subproduct_flag"`
	Suboptions SubOption `json:"suboptions"`
}

type SubOption struct {
	Options map[string]SubProductOption `json:"options"`
}

type SubProductOption struct {
	Key   string       `json:"key"`
	Name  string       `json:"name"`
	Items []SubProduct `json:"items"`
}

type SubProduct struct {
	Key        string   `json:"key"`
	Text       string   `json:"text"`
	ProductIds []string `json:"product_ids"`
}

func (pd ProductDetail) ToProduct(publisherID uuid.UUID) (*product.Product, error) {
	return product.NewProduct(product.NewProductArgs{
		Name:        pd.ProductName,
		ExternalID:  pd.ProductId,
		Price:       pd.ProductMaxPrice,
		Discounted:  pd.isDiscounted(),
		PublisherID: publisherID,
	})
}

func (pd ProductDetail) isDiscounted() bool {
	if pd.ProductSalePercent > 0 {
		return true
	}

	return pd.ProductMaxPrice != pd.ProductFullPrice
}
