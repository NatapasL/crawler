package productdetail

import (
	"manga-crawler/internal/domain/catalogupdate/product"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/domain/scraper"

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

func (pd ProductDetail) ToProduct(
	seriesFinder series.SeriesNameFinder,
	nameCleaner scraper.NameCleaner,
) (wholesale.Product, error) {
	publisherId, _ := uuid.Parse("01eb250e-57ad-4be1-8906-dc1527de6238")
	product, err := product.NewProduct(product.NewProductArgs{
		Name:       pd.ProductId,
		ExternalID: pd.ProductId,
		Price:      pd.ProductMaxPrice,
		Discounted: pd.isDiscounted(),
	})
	if err != nil {
		return nil, err
	}

	series, err := seriesFinder.MatchOrCreateSeriesByName(nameCleaner.Clean(pd.ProductName), publisherId)
	if err != nil {
		return nil, err
	}

	product.SetSeriesID(series.ID())
	return product, nil
}

func (pd ProductDetail) isDiscounted() bool {
	if pd.ProductSalePercent > 0 {
		return true
	}

	return pd.ProductMaxPrice != pd.ProductFullPrice
}
