dcl:
	air -c crawler.air.toml
rcl:
	go run ./cmd/crawler/
migrate:
	go run ./cmd/migration/