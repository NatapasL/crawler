package subproduct

import (
	"manga-crawler/internal/domain/scraper/productupdater"
	"manga-crawler/internal/domain/scraper/siamintershop"

	"github.com/google/uuid"
)

type SubProduct struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductTitle string `json:"product_title"`

	ProductMinPrice  float64 `json:"product_min_price"`
	ProductMaxPrice  float64 `json:"product_max_price"`
	ProductFullPrice float64 `json:"product_full_price"`

	ProductSalePercent float64 `json:"product_sale_percent"`

	ProductFullUrl string `json:"product_full_url"`
}

func (sp SubProduct) ToProduct() productupdater.Product {
	wholesaleID, _ := uuid.Parse(siamintershop.SiamintershopID)

	return productupdater.NewProduct(productupdater.NewProductArgs{
		Name:        sp.ProductName,
		WholesaleID: wholesaleID,
		ExternalID:  sp.ProductId,
		SourceUrl:   sp.ProductFullUrl,
		Price:       sp.ProductMaxPrice,
		Discounted:  sp.isDiscounted(),
	})
}

func (sp SubProduct) isDiscounted() bool {
	return sp.ProductMaxPrice != sp.ProductFullPrice
}
