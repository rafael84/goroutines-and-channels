SERVICE := simulation

.PHONY: $(SERVICE)

all: $(SERVICE)

$(SERVICE): ./bin
	@go build -o bin/$(SERVICE) cmd/$(SERVICE)/*.go

run-$(SERVICE): service
	@bin/$(SERVICE)

lint:
	@golint ./...

./bin:
	@mkdir -p ./bin
