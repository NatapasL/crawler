package siamintershop

import (
	"manga-crawler/internal/infrastructure/model"
	"time"

	"github.com/google/uuid"
)

type SeriesUpdater struct {
	nameCleaner      NameCleaner
	seriesRepository SeriesRepository
}

func NewSeriesUpdater(nameCleaner NameCleaner, seriesRepository SeriesRepository) *SeriesUpdater {
	return &SeriesUpdater{nameCleaner, seriesRepository}
}

func (su SeriesUpdater) UpdateSeries(products []SiamintershopProduct) {
	for _, product := range products {
		cleanedName := su.nameCleaner.Clean(product.ProductName)

		existingSeries := su.seriesRepository.FindByName(cleanedName)
		if len(existingSeries) < 1 {
			series := su.mapSiamintershopProductToSeries(product)
			su.seriesRepository.Persist(series)
		}
	}
}

func (su SeriesUpdater) mapSiamintershopProductToSeries(product SiamintershopProduct) model.Series {
	publisherId, _ := uuid.FromBytes([]byte("C941BB52-50BA-4273-81D6-2A56831B5B3A"))

	name := model.SeriesName{
		ID:          uuid.New(),
		Name:        su.nameCleaner.Clean(product.ProductName),
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
