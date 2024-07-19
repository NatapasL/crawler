dcl:
	air -c crawler.air.toml
rcl:
	go run ./cmd/crawler/main.go
migrate:
	go run ./cmd/migration/main.go