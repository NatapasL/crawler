package subproduct

import (
	"manga-crawler/internal/domain/catalogupdate/product"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/domain/scraper"

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

func (sp SubProduct) ToProduct(
	seriesFinder series.SeriesNameFinder,
	nameCleaner scraper.NameCleaner,
) (wholesale.Product, error) {
	publisherId, _ := uuid.Parse("01eb250e-57ad-4be1-8906-dc1527de6238")
	product, err := product.NewProduct(product.NewProductArgs{
		Name:       sp.ProductId,
		ExternalID: sp.ProductId,
		Price:      sp.ProductMaxPrice,
		Discounted: sp.isDiscounted(),
	})
	if err != nil {
		return nil, err
	}

	series, err := seriesFinder.MatchOrCreateSeriesByName(nameCleaner.Clean(sp.ProductName), publisherId)
	if err != nil {
		return nil, err
	}

	product.SetSeriesID(series.ID())
	return product, nil
}

func (sp SubProduct) isDiscounted() bool {
	if sp.ProductSalePercent > 0 {
		return true
	}

	return sp.ProductMaxPrice != sp.ProductFullPrice
}
