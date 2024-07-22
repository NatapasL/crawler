package productupdater

import (
	"manga-crawler/internal/domain/scraper/namecleaner"

	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	name        string
	wholesaleID uuid.UUID
	externalID  string
	sourceUrl   string
	seriesID    uuid.UUID

	Price ProductPrice
}

type NewProductArgs struct {
	Name        string
	WholesaleID uuid.UUID
	ExternalID  string
	SourceUrl   string
	Price       float64
	Discounted  bool
}

func NewProduct(args NewProductArgs) Product {
	product := Product{
		id:          uuid.New(),
		name:        args.Name,
		wholesaleID: args.WholesaleID,
		externalID:  args.ExternalID,
		sourceUrl:   args.SourceUrl,
		Price:       ProductPrice{price: args.Price, discounted: args.Discounted},
	}

	return product
}

func (p Product) ID() uuid.UUID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) WholesaleID() uuid.UUID {
	return p.wholesaleID
}

func (p Product) ToSeriesNameMatcher(nameCleaner namecleaner.NameCleaner) SeriesNameMatcher {
	return NewSeriesNameMatcher(NewSeriesNameMatcherArgs{
		Name: nameCleaner.Clean(p.name),
	})
}

func (p *Product) AttachToSeries(series Series) {
	p.seriesID = series.ID()
}
