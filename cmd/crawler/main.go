package main

import (
	"log"
	"manga-crawler/cmd/crawler/container"
	"manga-crawler/config"
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
	seriesNameMatcherModule := container.InitializeSeriesNameMatcherModule(container.SeriesNameMatcherModuleDependencies{
		RepositoryModule: repositoryModule,
	})
	siamintershopApiModule := container.InitializeSiamintershopApiModule(container.SiamintershopApiModuleDependency{})
	siamintershopModule := container.InitializeSiamintershopModule(container.SiamintershopModuleDependencies{
		ApiModule:               siamintershopApiModule,
		SeriesNameMatcherModule: seriesNameMatcherModule,
	})

	siamintershopModule.Scraper.ScrapeAll()
}
