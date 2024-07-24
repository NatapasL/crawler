package main

import (
	"log"
	"manga-crawler/cmd/crawler/container"
	"manga-crawler/cmd/crawler/container/seriesnamematcher"
	"manga-crawler/cmd/crawler/container/siamintershop"
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
	seriesNameMatcherModule := seriesnamematcher.InitializeSeriesNameMatcherModule(
		seriesnamematcher.SeriesNameMatcherModuleDependencies{
			RepositoryModule: repositoryModule,
		},
	)
	siamintershopModule := siamintershop.InitializeSiamintershopModule(siamintershop.SiamintershopModuleDependencies{
		SeriesNameMatcherModule: seriesNameMatcherModule,
	})
	log.Println(siamintershopModule)
}
