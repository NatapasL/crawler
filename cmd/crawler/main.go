package main

import (
	"log"
	"manga-crawler/config"
	"manga-crawler/internal/infrastructure/database/postgres"
	postgresrepository "manga-crawler/internal/infrastructure/repository/postgres_repository"
)

func main() {
	conf := config.GetConfig()
	db := postgres.GetConnection(conf.Postgres)
	if db == nil {
		log.Println("Cant connect to db...")
		return
	}

	seriesRepository := postgresrepository.NewSeriesRepository(db)
	seriesNameMatcherRepository := postgresrepository.NewSeriesNameMatcherRepository(db)
	siamintershopContainer := NewSiamintershopScraperContainer(seriesRepository, seriesNameMatcherRepository)

	siamintershopContainer.scraper.Scrape()
}
