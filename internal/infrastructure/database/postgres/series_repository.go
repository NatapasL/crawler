package postgres

import (
	"fmt"
	"log"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriesRepository struct {
	db *gorm.DB
}

func NewSeriesRepository(db *gorm.DB) *SeriesRepository {
	return &SeriesRepository{db}
}

func (repository SeriesRepository) FindByName(name string) []model.Series {
	var seriesNames []model.SeriesName
	repository.db.Select("series_id").Distinct("series_id").Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&seriesNames)

	var series []model.Series

	if len(seriesNames) < 1 {
		return series
	}

	var seriesIds []uuid.UUID
	for _, seriesName := range seriesNames {
		seriesIds = append(seriesIds, seriesName.SeriesID)
	}

	repository.db.Joins("LEFT JOIN series_name ON series_name.series_id = series.id").Where("series.id IN ?", seriesIds).Find(&series)

	return series
}

func (repository SeriesRepository) Persist(series model.Series) (*uuid.UUID, error) {
	createdSeries := model.Series{
		ID:          series.ID,
		PublisherID: series.PublisherID,
		CreatedAt:   series.CreatedAt,
		UpdatedAt:   series.UpdatedAt,
	}

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(&createdSeries)

		for _, name := range series.Names {
			tx.Save(&model.SeriesName{
				ID:          name.ID,
				SeriesID:    createdSeries.ID,
				Name:        name.Name,
				Language:    name.Language,
				DefaultName: name.DefaultName,
			})
		}

		return nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &createdSeries.ID, nil

}
