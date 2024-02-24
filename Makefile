build: modules
	@mkdir -p bin
	go build -o ./bin/goblog ./...

lint: modules
	@golangci-lint run

modules:
	@go mod tidy

clean:
	@rm -rf bin
