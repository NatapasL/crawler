package siamintershop

import "manga-crawler/internal/infrastructure/model"

type SeriesRepository interface {
	FindByName(string) []model.Series
}
