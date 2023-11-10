.PHONY: run redis

run: redis
	@go run main.go

redis:
	@docker-compose up -d
