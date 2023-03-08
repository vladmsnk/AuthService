compose-up:
	docker-compose up --build -d
.PHONY: compose-up

compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

docker-rm-volume:
	docker volume rm pg-data
.PHONY: docker-rm-volume

test:
	go test -v -cover -race ./internal/...
.PHONY: test

linter-golangci:
	golangci-lint run
.PHONY: linter-golangci

migrate-up:
	migrate -path migrations -database '$(PG_URL)' up
.PHONY: migrate-up