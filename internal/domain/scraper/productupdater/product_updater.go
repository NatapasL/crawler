package productupdater

import (
	"manga-crawler/internal/domain/scraper/namecleaner"

	"github.com/google/uuid"
)

type ProductUpdater struct {
	productRepository          ProductRepository
	nameCleaner                namecleaner.NameCleaner
	seriesFinderService        SeriesFinderService
	seriesRepository           SeriesRepository
	nameMatcherAttacherService NameMatcherAttacherService
}

type ProductUpdaterDependencies struct {
	ProductRepository          ProductRepository
	NameCleaner                namecleaner.NameCleaner
	SeriesFinderService        SeriesFinderService
	SeriesRepository           SeriesRepository
	NameMatcherAttacherService NameMatcherAttacherService
}

func NewProductUpdater(deps ProductUpdaterDependencies) *ProductUpdater {
	return &ProductUpdater{
		productRepository:          deps.ProductRepository,
		nameCleaner:                deps.NameCleaner,
		seriesFinderService:        deps.SeriesFinderService,
		seriesRepository:           deps.SeriesRepository,
		nameMatcherAttacherService: deps.NameMatcherAttacherService,
	}
}

func (pu ProductUpdater) UpdateMultiple(products []Product, publisherId uuid.UUID) {
	for _, product := range products {
		pu.Update(product, publisherId)
	}
}

func (pu ProductUpdater) Update(product Product, publisherId uuid.UUID) error {
	existingProduct, err := pu.productRepository.FindByWholesaleIDAndExternalID(product.WholesaleID(), product.externalID)
	if err != nil {
		return err
	}

	if existingProduct != nil {
		return pu.updateExistingProduct(*existingProduct, product)
	}

	return pu.addNewProduct(product, publisherId)
}

func (pu ProductUpdater) updateExistingProduct(existingProduct Product, newProduct Product) error {
	existingProduct.Price = newProduct.Price
	_, err := pu.productRepository.Persist(existingProduct)

	return err
}

func (pu ProductUpdater) addNewProduct(product Product, publisherId uuid.UUID) error {
	matcher := product.ToSeriesNameMatcher(pu.nameCleaner)

	var series *Series
	series, err := pu.seriesFinderService.FindByNameMatcher(matcher)
	if err != nil {
		return err
	}

	if series == nil {
		series, err = pu.nameMatcherAttacherService.AttachToNonExistingSeries(matcher, publisherId)
		if err != nil {
			return err
		}
	}

	product.AttachToSeries(*series)
	_, err = pu.productRepository.Persist(product)
	return err
}
