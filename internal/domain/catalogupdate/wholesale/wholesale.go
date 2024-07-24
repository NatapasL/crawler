package wholesale

import "github.com/google/uuid"

type Wholesale struct {
	id       uuid.UUID
	products []Product
}

func (wh *Wholesale) UpdateCatalog(s Scraper, next <-chan bool, productAdded chan<- bool) error {
	productChannel := make(chan Product)
	go func() {
		for {
			product, ok := <-productChannel
			if !ok {
				return
			}

			wh.clearProducts()
			wh.addProduct(product)
			productAdded <- true
		}
	}()

	err := s.Scrape(next, productChannel)
	return err
}

func (wh *Wholesale) addProduct(p Product) {
	p.SetWholesaleID(wh.id)
	wh.products = append(wh.products, p)
}

func (wh *Wholesale) clearProducts() {
	wh.products = []Product{}
}

func (wh Wholesale) ID() uuid.UUID {
	return wh.id
}
