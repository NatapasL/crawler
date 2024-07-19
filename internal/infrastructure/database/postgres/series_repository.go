package postgres

import (
	"fmt"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriesRepository struct {
	db *gorm.DB
}

func (repository SeriesRepository) FindByName(name string) []model.Series {
	var seriesNames []model.SeriesName
	repository.db.Select("series_id").Distinct("series_id").Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&seriesNames)

	var seriesIds []uuid.UUID
	for _, seriesName := range seriesNames {
		seriesIds = append(seriesIds, seriesName.SeriesID)
	}

	var series []model.Series
	repository.db.Joins("series_name").Where("id IN ?", seriesIds).Find(&series)

	return series
}
