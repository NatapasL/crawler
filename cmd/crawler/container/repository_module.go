package container

import (
	postgresrepository "manga-crawler/internal/infrastructure/repository/postgres"

	"gorm.io/gorm"
)

type RepositoryModule struct {
	SeriesRepository            *postgresrepository.SeriesRepository
	SeriesNameMatcherRepository *postgresrepository.SeriesNameMatcherRepository
}

type RepositoryModuleDependencies struct {
	DB *gorm.DB
}

func InitializeRepositoryModule(deps RepositoryModuleDependencies) RepositoryModule {
	seriesRepository := postgresrepository.NewSeriesRepository(deps.DB)
	seriesNameMatcherRepository := postgresrepository.NewSeriesNameMatcherRepository(deps.DB)

	return RepositoryModule{
		SeriesRepository:            seriesRepository,
		SeriesNameMatcherRepository: seriesNameMatcherRepository,
	}
}
