package updatecatalog

import "manga-crawler/internal/domain/catalogupdate/product"

type productRepository interface {
	Persist(product.Product)
}
