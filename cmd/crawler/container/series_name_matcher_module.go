package container

import (
	seriesnamematcher "manga-crawler/internal/domain/scraper/series_name_matcher"
)

type SeriesNameMatcherModule struct {
	SeriesNameMatcherToSeriesNameMapper *seriesnamematcher.SeriesNameMatcherToSeriesNameMapper
	SeriesNameMatcherToSeriesMapper     *seriesnamematcher.SeriesNameMatcherToSeriesMapper
	SeriesNameMatcherService            *seriesnamematcher.SeriesNameMatcherService
}

type SeriesNameMatcherModuleDependencies struct {
	RepositoryModule RepositoryModule
}

func InitializeSeriesNameMatcherModule(deps SeriesNameMatcherModuleDependencies) SeriesNameMatcherModule {
	seriesNameMatcherToSeriesNameMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesNameMapper()
	seriesNameMatcherToSeriesMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesMapper(*seriesNameMatcherToSeriesNameMapper)
	seriesNameMatcherService := seriesnamematcher.NewSeriesNameMatcherService(
		deps.RepositoryModule.SeriesNameMatcherRepository,
		deps.RepositoryModule.SeriesRepository,
		*seriesNameMatcherToSeriesMapper,
	)

	return SeriesNameMatcherModule{
		SeriesNameMatcherToSeriesNameMapper: seriesNameMatcherToSeriesNameMapper,
		SeriesNameMatcherToSeriesMapper:     seriesNameMatcherToSeriesMapper,
		SeriesNameMatcherService:            seriesNameMatcherService,
	}
}
