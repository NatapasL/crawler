package postgres

import (
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

func (repository SeriesNameMatcherRepository) FindByName(name string) []model.SeriesNameMatcher {
	var seriesNameMatchers []model.SeriesNameMatcher

	repository.db.Distinct("series_id").Where("name = ?", name).Find(&seriesNameMatchers)

	return seriesNameMatchers
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
