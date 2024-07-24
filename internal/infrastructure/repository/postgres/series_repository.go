package postgres

import (
	"log"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/language"
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

func (repository SeriesRepository) FindById(id uuid.UUID) (*series.Series, error) {
	var s model.Series
	repository.db.Joins(
		"series_name ON series_name.series_id = series.id",
	).Joins(
		"series_name_matcher ON series_name_matcher.series_id = series.id",
	).Where(
		"series.id = ?", id,
	).Find(&s)

	var names []series.ExistingSeriesNameArgs
	for _, name := range s.Names {
		names = append(names, series.ExistingSeriesNameArgs{
			ID:        name.ID,
			Name:      name.Name,
			Language:  language.Language(name.Language),
			IsDefault: name.DefaultName,
		})
	}

	var nameMatchers []series.ExistingSeriesNameMatcherArgs
	for _, nameMatcher := range s.NameMatchers {
		nameMatchers = append(nameMatchers, series.ExistingSeriesNameMatcherArgs{
			ID:       nameMatcher.ID,
			Name:     nameMatcher.Name,
			SeriesID: nameMatcher.SeriesID,
		})
	}

	return series.ExistingSeries(series.ExistingSeriesArgs{
		ID:           s.ID,
		PublisherID:  s.PublisherID,
		NameMatchers: nameMatchers,
		Names:        names,
	})
}

func (repository SeriesRepository) Persist(series series.Series) (*uuid.UUID, error) {
	savedSeries := model.Series{
		ID:          series.ID(),
		PublisherID: series.PublisherID(),
	}

	seriesNames := series.Names()
	var savedNames []model.SeriesName
	for _, seriesName := range seriesNames {
		savedNames = append(savedNames, model.SeriesName{
			ID:          seriesName.ID(),
			SeriesID:    series.ID(),
			Name:        seriesName.Name(),
			Language:    seriesName.Name(),
			DefaultName: seriesName.IsDefault(),
		})
	}

	nameMatchers := series.NameMatchers()
	var savedNameMatchers []model.SeriesNameMatcher
	for _, nameMatcher := range nameMatchers {
		savedNameMatchers = append(savedNameMatchers, model.SeriesNameMatcher{
			ID:       nameMatcher.ID(),
			SeriesID: series.ID(),
			Name:     nameMatcher.Name(),
		})
	}

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&savedSeries).Error
		if err != nil {
			return err
		}

		err = tx.Save(savedNames).Error
		if err != nil {
			return err
		}

		err = tx.Save(savedNameMatchers).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &savedSeries.ID, nil

}
