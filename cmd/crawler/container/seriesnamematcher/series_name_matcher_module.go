package seriesnamematcher

import (
	"manga-crawler/cmd/crawler/container"
	"manga-crawler/internal/domain/seriesnamematcher"
)

type SeriesNameMatcherModule struct {
	SeriesNameMatcherToSeriesNameMapper *seriesnamematcher.SeriesNameMatcherToSeriesNameMapper
	SeriesNameMatcherToSeriesMapper     *seriesnamematcher.SeriesNameMatcherToSeriesMapper
	SeriesNameMatcherService            *seriesnamematcher.SeriesNameMatcherService
}

type SeriesNameMatcherModuleDependencies struct {
	RepositoryModule container.RepositoryModule
}

func InitializeSeriesNameMatcherModule(deps SeriesNameMatcherModuleDependencies) SeriesNameMatcherModule {
	seriesNameMatcherToSeriesNameMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesNameMapper()
	seriesNameMatcherToSeriesMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesMapper(
		seriesnamematcher.SeriesNameMatcherToSeriesMapperDependencies{
			SeriesNameMatcherToSeriesNameMapper: *seriesNameMatcherToSeriesNameMapper,
		},
	)
	seriesNameMatcherService := seriesnamematcher.NewSeriesNameMatcherService(
		seriesnamematcher.SeriesNameMatcherServiceDependencies{
			SeriesNameMatcherRepository:     deps.RepositoryModule.SeriesNameMatcherRepository,
			SeriesRepository:                deps.RepositoryModule.SeriesRepository,
			SeriesNameMatcherToSeriesMapper: *seriesNameMatcherToSeriesMapper,
		},
	)

	return SeriesNameMatcherModule{
		SeriesNameMatcherToSeriesNameMapper: seriesNameMatcherToSeriesNameMapper,
		SeriesNameMatcherToSeriesMapper:     seriesNameMatcherToSeriesMapper,
		SeriesNameMatcherService:            seriesNameMatcherService,
	}
}
