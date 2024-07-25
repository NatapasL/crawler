package postgres

import (
	"errors"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriesNameMatcherRepository struct {
	db *gorm.DB
}

func NewSeriesNameMatcherRepository(db *gorm.DB) *SeriesNameMatcherRepository {
	return &SeriesNameMatcherRepository{db}
}

func (repository SeriesNameMatcherRepository) FindByName(name string) (*series.SeriesNameMatcher, error) {
	var seriesNameMatchers model.SeriesNameMatcher
	err := repository.db.Where("name = ?", name).First(&seriesNameMatchers).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return series.ExistingSeriesNameMatcher(series.ExistingSeriesNameMatcherArgs{
		ID:       seriesNameMatchers.ID,
		Name:     seriesNameMatchers.Name,
		SeriesID: seriesNameMatchers.SeriesID,
	})
}

func (repository SeriesNameMatcherRepository) FindById(id uuid.UUID) (model.SeriesNameMatcher, error) {
	var seriesNameMatcher model.SeriesNameMatcher
	if err := repository.db.Where("id = ?", id).First(&seriesNameMatcher).Error; err != nil {
		return seriesNameMatcher, err
	}

	return seriesNameMatcher, nil
}

func (repository SeriesNameMatcherRepository) Persist(nameMatcher model.SeriesNameMatcher) uuid.UUID {
	createdNameMatcher := model.SeriesNameMatcher{
		ID:       nameMatcher.ID,
		SeriesID: nameMatcher.SeriesID,
		Name:     nameMatcher.Name,
	}
	repository.db.Save(&createdNameMatcher)

	return createdNameMatcher.ID
}
