package productsearch

import (
	"manga-crawler/internal/domain/catalogupdate/product"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/domain/scraper"

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

func (p ProductSearchResponseProduct) ToProduct(
	seriesFinder series.SeriesNameFinder,
	nameCleaner scraper.NameCleaner,
	publisherID uuid.UUID,
) (wholesale.Product, error) {
	product, err := product.NewProduct(product.NewProductArgs{
		Name:       p.ProductName,
		ExternalID: p.ProductId,
		Price:      p.ProductMaxPrice,
		Discounted: p.isDiscounted(),
	})
	if err != nil {
		return nil, err
	}

	series, err := seriesFinder.MatchOrCreateSeriesByName(nameCleaner.Clean(p.ProductName), publisherID)
	if err != nil {
		return nil, err
	}

	product.SetSeriesID(series.ID())
	return product, nil
}

func (p ProductSearchResponseProduct) isDiscounted() bool {
	if p.ProductSalePercent > 0 {
		return true
	}

	return p.ProductMaxPrice != p.ProductFullPrice
}
