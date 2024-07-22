package siamintershopproductsearch

import "time"

type productSearchScraperIntervalConfig struct {
	chunkSize int
	delay     time.Duration
}

func defaultIntervalConfig() productSearchScraperIntervalConfig {
	return productSearchScraperIntervalConfig{
		chunkSize: 59,
		delay:     1 * time.Second,
	}
}
