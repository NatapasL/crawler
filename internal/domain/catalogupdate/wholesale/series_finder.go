package wholesale

type SeriesFinder interface {
	FindSeriesForProduct(product Product) (Series, error)
}
