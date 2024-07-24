package wholesale

import "github.com/google/uuid"

type Wholesale struct {
	id       uuid.UUID
	products []Product
}

type ExistingWholesaleArgs struct {
	ID uuid.UUID
}

func ExistingWholesale(args ExistingWholesaleArgs) (*Wholesale, error) {
	ws := Wholesale{id: args.ID}
	err := ws.validate()
	if err != nil {
		return nil, err
	}

	return &ws, nil
}

func (Wholesale) validate() error {
	return nil
}

func (ws *Wholesale) UpdateCatalog(s Scraper, next <-chan bool, productAdded chan<- bool) error {
	defer close(productAdded)

	productChannel := make(chan Product)
	go func() {
		for {
			product, ok := <-productChannel
			if !ok {
				return
			}

			ws.clearProducts()
			ws.addProduct(product)
			productAdded <- true
		}
	}()

	err := s.Scrape(next, productChannel)
	return err
}

func (ws *Wholesale) addProduct(p Product) {
	p.SetWholesaleID(ws.id)
	ws.products = append(ws.products, p)
}

func (ws *Wholesale) clearProducts() {
	ws.products = []Product{}
}

func (ws Wholesale) ID() uuid.UUID {
	return ws.id
}

func (ws Wholesale) Products() []Product {
	return ws.products
}
