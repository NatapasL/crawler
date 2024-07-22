package siamintershopproductdetail

import (
	"manga-crawler/internal/domain/scraper/namecleaner"
	seriesnamematcher "manga-crawler/internal/domain/scraper/series_name_matcher"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type ResponseService struct {
	nameCleaner              namecleaner.NameCleaner
	seriesNameMatcherService seriesnamematcher.SeriesNameMatcherService
}

type ResponseServiceDependencies struct {
	NameCleaner              namecleaner.NameCleaner
	SeriesNameMatcherService seriesnamematcher.SeriesNameMatcherService
}

func NewResponseService(deps ResponseServiceDependencies) *ResponseService {
	return &ResponseService{nameCleaner: deps.NameCleaner, seriesNameMatcherService: deps.SeriesNameMatcherService}
}

func (service ResponseService) AddSeriesNameMatcherIfNotExists(response ProductDetailResponse) error {
	nameMatcher := service.mapResponseProductToSeriesNameMatcher(response)

	publisherId, _ := uuid.FromBytes([]byte("01eb250e-57ad-4be1-8906-dc1527de6238"))
	err := service.seriesNameMatcherService.CreateIfNotExists(nameMatcher, publisherId)
	if err != nil {
		return err
	}

	return nil
}

func (service ResponseService) mapResponseProductToSeriesNameMatcher(product ProductDetailResponse) model.SeriesNameMatcher {
	return model.SeriesNameMatcher{
		ID:   uuid.New(),
		Name: service.nameCleaner.Clean(product.ProductName),
	}
}
