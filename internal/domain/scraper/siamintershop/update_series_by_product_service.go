package siamintershop

import (
	siamintershoprepository "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_repository"
	"manga-crawler/internal/infrastructure/model"
	"time"

	"github.com/google/uuid"
)

type UpdateSeriesByProductService struct {
	nameCleaner                 NameCleaner
	seriesRepository            siamintershoprepository.SeriesRepository
	seriesNameMatcherRepository siamintershoprepository.SeriesNameMatcherRepository
}

func NewUpdateSeriesByProductService(
	nameCleaner NameCleaner,
	seriesRepository siamintershoprepository.SeriesRepository,
	seriesNameMatcherRepository siamintershoprepository.SeriesNameMatcherRepository,
) *UpdateSeriesByProductService {
	return &UpdateSeriesByProductService{nameCleaner, seriesRepository, seriesNameMatcherRepository}
}

func (su UpdateSeriesByProductService) UpdateSeries(products []SiamintershopProduct) []error {
	var errs []error
	for _, product := range products {
		isExists := su.isExists(product)

		if isExists {
			continue
		}

		nameMatcherId := su.addNameMatcher(product)

		seriesId, err := su.addSeries(product)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		su.attachSeriesToNameMatcher(*seriesId, nameMatcherId)
	}

	return errs
}

func (su UpdateSeriesByProductService) isExists(product SiamintershopProduct) bool {
	cleanedName := su.cleanedProductName(product)

	nameMatchers := su.seriesNameMatcherRepository.FindByName(cleanedName)

	return len(nameMatchers) > 0
}

func (su UpdateSeriesByProductService) addSeries(product SiamintershopProduct) (*uuid.UUID, error) {
	series := su.mapSiamintershopProductToSeries(product)
	seriesId, err := su.seriesRepository.Persist(series)
	if err != nil {
		return nil, err
	}

	return seriesId, nil
}

func (su UpdateSeriesByProductService) mapSiamintershopProductToSeries(product SiamintershopProduct) model.Series {
	publisherId, _ := uuid.FromBytes([]byte("C941BB52-50BA-4273-81D6-2A56831B5B3A"))

	name := model.SeriesName{
		ID:          uuid.New(),
		Name:        su.cleanedProductName(product),
		Language:    "TH",
		DefaultName: false,
	}

	series := model.Series{
		ID:          uuid.New(),
		PublisherID: publisherId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Names:       []model.SeriesName{name},
	}

	return series
}

func (su UpdateSeriesByProductService) addNameMatcher(product SiamintershopProduct) uuid.UUID {
	nameMatcher := model.SeriesNameMatcher{
		ID:   uuid.New(),
		Name: su.cleanedProductName(product),
	}

	return su.seriesNameMatcherRepository.Persist(nameMatcher)
}

func (su UpdateSeriesByProductService) cleanedProductName(product SiamintershopProduct) string {
	return su.nameCleaner.Clean(product.ProductName)
}

func (su UpdateSeriesByProductService) attachSeriesToNameMatcher(seriesId uuid.UUID, nameMatcherId uuid.UUID) error {
	nameMatcher, err := su.seriesNameMatcherRepository.FindById(nameMatcherId)
	if err != nil {
		return err
	}

	nameMatcher.SetSeriesID(seriesId)
	su.seriesNameMatcherRepository.Persist(nameMatcher)
	return nil
}
