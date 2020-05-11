CLIENT := client-service
SERVER := server-service

.PHONY: $(CLIENT) $(SERVER)

all: $(CLIENT) $(SERVER)

$(CLIENT): ./bin
	@go build -o bin/$(CLIENT) cmd/$(CLIENT)/*.go

$(SERVER): ./bin
	@go build -o bin/$(SERVER) cmd/$(SERVER)/*.go

run-$(CLIENT): $(CLIENT)
	@bin/$(CLIENT)

run-$(SERVER): $(SERVER)
	@bin/$(SERVER)

lint:
	@golint ./...

./bin:
	@mkdir -p ./bin
