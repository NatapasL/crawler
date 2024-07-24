package product

import "github.com/google/uuid"

type Product struct {
	id          uuid.UUID
	name        string
	seriesID    uuid.UUID
	wholesaleID uuid.UUID
	externalID  string
	price       float64
	discounted  bool
}

type NewProductArgs struct {
	Name       string
	SeriesID   uuid.UUID
	ExternalID string
	Price      float64
	Discounted bool
}

func NewProduct(args NewProductArgs) (*Product, error) {
	product := Product{
		id:         uuid.New(),
		seriesID:   args.SeriesID,
		name:       args.Name,
		externalID: args.ExternalID,
		price:      args.Price,
		discounted: args.Discounted,
	}
	err := product.validate()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) SetWholesaleID(wholesaleID uuid.UUID) {
	p.wholesaleID = wholesaleID
}

func (Product) validate() error {
	return nil
}
