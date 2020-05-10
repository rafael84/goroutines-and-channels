CLIENT := client-service
SERVER := server-service

.PHONY: client server

all: client server

client: ./bin
	@go build -o bin/$(CLIENT) cmd/$(CLIENT)/*.go

server: ./bin
	@go build -o bin/$(SERVER) cmd/$(SERVER)/*.go

run-client: client
	@bin/$(CLIENT)

run-server: server
	@bin/$(SERVER)

lint:
	@golint ./...

./bin:
	@mkdir -p ./bin
