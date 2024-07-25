package product

import (
	"github.com/google/uuid"
)

type Product struct {
	productID   uuid.UUID
	priceID     uuid.UUID
	name        string
	seriesID    uuid.UUID
	wholesaleID uuid.UUID
	externalID  string
	price       ProductPrice
	publisherID uuid.UUID
}

type NewProductArgs struct {
	Name        string
	ExternalID  string
	Price       float64
	Discounted  bool
	PublisherID uuid.UUID
}

func NewProduct(args NewProductArgs) (*Product, error) {
	price, err := newProductPrice(newProductPriceArgs{price: args.Price, discounted: args.Discounted})
	if err != nil {
		return nil, err
	}

	product := Product{
		productID:   uuid.New(),
		priceID:     uuid.New(),
		name:        args.Name,
		externalID:  args.ExternalID,
		price:       *price,
		publisherID: args.PublisherID,
	}
	err = product.validate()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) SetWholesaleID(wholesaleID uuid.UUID) {
	p.wholesaleID = wholesaleID
}

func (p *Product) SetSeriesID(seriesID uuid.UUID) {
	p.seriesID = seriesID
}

func (Product) validate() error {
	return nil
}

func (p Product) ProductID() uuid.UUID {
	return p.productID
}

func (p Product) Name() string {
	return p.name
}

func (p Product) SeriesID() uuid.UUID {
	return p.seriesID
}

func (p Product) ExternalID() string {
	return p.externalID
}

func (p Product) Price() float64 {
	return p.price.price
}

func (p Product) Discounted() bool {
	return p.price.discounted
}

func (p Product) PriceID() uuid.UUID {
	return p.priceID
}

func (p Product) WholesaleID() uuid.UUID {
	return p.wholesaleID
}

func (p Product) PublisherID() uuid.UUID {
	return p.publisherID
}
