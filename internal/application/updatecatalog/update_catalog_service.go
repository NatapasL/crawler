package updatecatalog

import (
	"manga-crawler/internal/domain/catalogupdate/product"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/scraper"
	"manga-crawler/internal/domain/scraper/namecleaner"

	"github.com/google/uuid"
)

type UpdateCatalogService struct {
	seriesFinder        *series.SeriesFinderService
	wholesaleRepository wholesaleRepository
	productRepository   productRepository
	scraperFactory      *scraper.ScraperFactory
	nameCleanerFactory  *namecleaner.NameCleanerFactory
}

type UpdateCatalogServiceDependencies struct {
	SeriesFinder        *series.SeriesFinderService
	WholesaleRepository wholesaleRepository
	ProductRepository   productRepository
	ScraperFactory      *scraper.ScraperFactory
	NameCleanerFactory  *namecleaner.NameCleanerFactory
}

func NewUpdateCatalogService(deps UpdateCatalogServiceDependencies) *UpdateCatalogService {
	return &UpdateCatalogService{
		seriesFinder:        deps.SeriesFinder,
		wholesaleRepository: deps.WholesaleRepository,
		productRepository:   deps.ProductRepository,
		scraperFactory:      deps.ScraperFactory,
		nameCleanerFactory:  deps.NameCleanerFactory,
	}
}

func (service UpdateCatalogService) UpdateWholesaleCatalog(wholesaleID uuid.UUID) error {
	ws, err := service.wholesaleRepository.FindById(wholesaleID)
	if err != nil {
		return err
	}

	scraper, err := service.scraperFactory.GetScraper(ws.ID().String())
	if err != nil {
		return err
	}

	product := make(chan product.Product)
	defer close(product)

	nameCleaner, err := service.nameCleanerFactory.GetCleaner(ws.ID().String())
	if err != nil {
		return err
	}

	go func() {
		for {
			p := <-product
			cleanedName := nameCleaner.Clean(p.Name())
			go func() {
				series, err := service.seriesFinder.MatchOrCreateSeriesByName(cleanedName, p.PublisherID())
				if err != nil {
					return
				}
				p.SetSeriesID(series.ID())
				p.SetWholesaleID(ws.ID())
				service.productRepository.Persist(p)
			}()
		}
	}()

	scraper.Scrape(product)

	return nil
}
