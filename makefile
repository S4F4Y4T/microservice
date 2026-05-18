.PHONY: run build dev clean test lint tidy \
        db-up db-down \
        migrate-up migrate-down migrate-force migrate-create

BIN           := ./bin/api
MIGRATION_DIR := database/migration
DB_URL        ?= postgres://postgres:postgres@localhost:5433/rest_db?sslmode=disable

run:
	go run ./cmd/api/main.go

build:
	@rm -rf $(BIN)
	@mkdir -p bin
	go build -o $(BIN) ./cmd/api/main.go

dev:
	air

clean:
	rm -rf bin tmp

test:
	go test ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

migrate-up:
	migrate -path database/migrations \
	-database "postgres://postgres:postgres@localhost:5433/rest_db?sslmode=disable" \
	up

migrate-down:
	migrate -path database/migrations \
	-database "postgres://postgres:postgres@localhost:5433/rest_db?sslmode=disable" \
	down

migrate-create:
	migrate create -ext sql -dir database/migrations -seq $(name)