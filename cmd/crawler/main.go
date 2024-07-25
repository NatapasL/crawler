package main

import (
	"log"
	"manga-crawler/cmd/crawler/container"
	"manga-crawler/cmd/crawler/container/siamintershop"
	"manga-crawler/config"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/lookup/wholesale"
	"manga-crawler/internal/domain/scraper"
	"manga-crawler/internal/infrastructure/database/postgres"
)

func main() {
	conf := config.GetConfig()
	db := postgres.GetConnection(conf.Postgres)
	if db == nil {
		log.Println("Cant connect to db...")
		return
	}

	repositoryModule := container.InitializeRepositoryModule(container.RepositoryModuleDependencies{DB: db})
	seriesFinderService := series.NewSeriesFinderService(series.NewSeriesFinderServiceDependencies{
		SeriesRepository:            repositoryModule.SeriesRepository,
		SeriesNameMatcherRepository: repositoryModule.SeriesNameMatcherRepository,
	})
	siamintershopModule := siamintershop.InitializeSiamintershopModule(siamintershop.SiamintershopModuleDependencies{
		SeriesFinder: *seriesFinderService,
	})

	scraperFactory := scraper.NewScraperFactory()
	scraperFactory.Register(string(wholesale.SiamintershopID), siamintershopModule.Scraper)
}
