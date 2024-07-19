package main

import (
	"log"
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

	postgres.Migrate(db)
}
