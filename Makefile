.PHONY: default run build test docs clean

APP_NAME=search-zip-code

default: run

run:
	@go run cmd/api/main.go

build:
	@go build -o $(APP_NAME) cmd/api/main.go

test:
	@cd internal/pkg && go test -v
	@cd internal/util && go test -v

docs:
	@swag init -g cmd/api/main.go

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs