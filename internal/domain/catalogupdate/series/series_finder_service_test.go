package series

import (
	"testing"

	"github.com/google/uuid"
)

type stubSeriesNameMatcherRepository struct {
	seriesNameMatchers []SeriesNameMatcher
}

func (r stubSeriesNameMatcherRepository) FindByName(name string) (*SeriesNameMatcher, error) {
	for _, snm := range r.seriesNameMatchers {
		if snm.name == name {
			return &snm, nil
		}
	}

	return nil, nil
}

type stubSeriesRepository struct {
	series          []Series
	persistedSeries *Series
}

func (r stubSeriesRepository) FindById(id uuid.UUID) (*Series, error) {
	for _, series := range r.series {
		if series.id == id {
			return &series, nil
		}
	}

	return nil, nil
}

func (r *stubSeriesRepository) Persist(s Series) error {
	r.persistedSeries = &s
	return nil
}

var publisherID = uuid.New()

func TestSeriesFinderService_MatchOrCreateSeriesByName(t *testing.T) {
	t.Run("Create new series when name matcher not found", func(t *testing.T) {
		// Arrange
		seriesNameMatcherRepository := stubSeriesNameMatcherRepository{}
		seriesRepository := stubSeriesRepository{}
		sut := NewSeriesFinderService(NewSeriesFinderServiceDependencies{
			SeriesRepository:            &seriesRepository,
			SeriesNameMatcherRepository: &seriesNameMatcherRepository,
		})

		// Act
		series, err := sut.MatchOrCreateSeriesByName("series_name", publisherID)

		// Assert
		if err != nil {
			t.Fatal("Error: ", err)
		}
		if series == nil {
			t.Fatal("No series return.")
		}
		if series.publisherID != publisherID {
			t.Fatalf("Created series publisherID incorrect, expect %s got %s", publisherID, series.publisherID)
		}
		if seriesRepository.persistedSeries == nil {
			t.Fatal("Series not persisted")
		}
		if seriesRepository.persistedSeries.id != series.id {
			t.Fatalf("Persisted series incorrect, expect id %s got %s", series.id, seriesRepository.persistedSeries.id)
		}
	})

	t.Run("Return series found", func(t *testing.T) {
		// Arrange
		seriesId := uuid.New()
		seriesId2 := uuid.New()
		seriesName := "series_name"

		seriesNameMatcherRepository := stubSeriesNameMatcherRepository{
			seriesNameMatchers: []SeriesNameMatcher{
				{seriesID: seriesId, name: seriesName},
				{seriesID: seriesId2, name: "other name"},
			},
		}
		seriesRepository := stubSeriesRepository{
			series: []Series{{id: seriesId}, {id: seriesId2}},
		}

		sut := NewSeriesFinderService(NewSeriesFinderServiceDependencies{
			SeriesRepository:            &seriesRepository,
			SeriesNameMatcherRepository: &seriesNameMatcherRepository,
		})

		// Act
		series, err := sut.MatchOrCreateSeriesByName(seriesName, publisherID)

		// Assert
		if err != nil {
			t.Fatal("Error: ", err)
		}
		if series == nil {
			t.Fatal("Series should be found")
		}
		if series.id != seriesId {
			t.Fatalf("Found incorrect series, expect id %s got %s", seriesId, series.id)
		}
	})

	t.Run("Return nil when series not found", func(t *testing.T) {
		// Arrange
		seriesName := "series_name"
		seriesNameMatcherRepository := stubSeriesNameMatcherRepository{
			seriesNameMatchers: []SeriesNameMatcher{{name: seriesName, seriesID: uuid.New()}},
		}
		seriesRepository := stubSeriesRepository{
			series: []Series{{id: uuid.New()}},
		}
		sut := NewSeriesFinderService(NewSeriesFinderServiceDependencies{
			SeriesRepository:            &seriesRepository,
			SeriesNameMatcherRepository: &seriesNameMatcherRepository,
		})

		// Act
		series, err := sut.MatchOrCreateSeriesByName(seriesName, publisherID)

		// Assert
		if err != nil {
			t.Error("Error: ", err)
		}
		if series != nil {
			t.Error("Series should not be found")
		}
	})
}
