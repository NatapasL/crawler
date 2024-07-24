package product

import "github.com/google/uuid"

type ProductPrice struct {
	id         uuid.UUID
	price      float64
	discounted bool
}

type newProductPriceArgs struct {
	price      float64
	discounted bool
}

func newProductPrice(args newProductPriceArgs) (*ProductPrice, error) {
	price := ProductPrice{id: uuid.New(), price: args.price, discounted: args.discounted}
	err := price.validate()
	if err != nil {
		return nil, err
	}

	return &price, nil
}

func (ProductPrice) validate() error {
	return nil
}
