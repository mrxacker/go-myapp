APP_NAME := myApp
BIN_DIR := bin
COMPOSE_FILE=docker-compose.yml
ENV_FILE=.env

.PHONY: run build test clean proto

proto:
	protoc --go_out=. --go_opt=paths=import \
	       --go-grpc_out=. --go-grpc_opt=paths=import \
	       proto/user/v1/user.proto

run: proto
	go run cmd/api/main.go

build: proto
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/ api/

docker-up:
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) up -d --build

docker-down:
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) down -v
