package productupdater

import "github.com/google/uuid"

type ProductPrice struct {
	id         uuid.UUID
	price      float64
	discounted bool
}

type newProductPriceArgs struct {
	Price      float64
	Discounted bool
}

func newProductPrice(args newProductPriceArgs) (*ProductPrice, error) {
	productPrice := ProductPrice{
		id:         uuid.New(),
		price:      args.Price,
		discounted: args.Discounted,
	}
	if err := productPrice.validate(); err != nil {
		return nil, err
	}

	return &productPrice, nil
}

func (p ProductPrice) validate() error {
	return nil
}
