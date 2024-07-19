package main

import (
	"log"
	"manga-crawler/internal/domain/scraper/siamintershop"
)

func main() {
	log.Println("hello")

	siamintershop.Scrape()
}
