package wholesale

import "github.com/google/uuid"

type Wholesale struct {
	id       uuid.UUID
	products []Product
}

func (wh *Wholesale) UpdateCatalog(s Scraper) error {
	products, err := s.Scrape()
	if err != nil {
		return err
	}
	for _, product := range products {
		wh.addProduct(product)
	}
	return nil
}

func (wh *Wholesale) addProduct(p Product) {
	p.SetWholesaleID(wh.id)
	wh.products = append(wh.products, p)
}

func (wh Wholesale) ID() uuid.UUID {
	return wh.id
}
