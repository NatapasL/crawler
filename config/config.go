package config

import "manga-crawler/internal/infrastructure/database/postgres"

const dbConnectionString = "host=localhost user=postgres password=1234 dbname=manga_crawler port=5432 sslmode=disable TimeZone=Asia/Bangkok"

type Config struct {
	Postgres postgres.PostgresConnectionConfig
	Http     HttpConfig
}

type HttpConfig struct {
	UserAgent string
}

var config *Config

func GetConfig() Config {
	if config == nil {
		config = &Config{
			Postgres: postgres.PostgresConnectionConfig{
				Host:     "localhost",
				User:     "postgres",
				Password: "1234",
				DbName:   "manga_crawler",
				Port:     "5432",
				SslMode:  "disable",
				TimeZone: "Asia/Bangkok",
			},
			Http: HttpConfig{
				UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
			},
		}
	}

	return *config
}
