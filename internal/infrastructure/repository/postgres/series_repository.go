package postgres

import (
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

func (repository SeriesRepository) FindByIds(ids []uuid.UUID) []model.Series {
	var series []model.Series
	repository.db.Joins("series_name ON series_name.series_id = series.id").Where("series.id IN ?", ids).Find(&series)

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
