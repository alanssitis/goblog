build: modules
	@mkdir -p bin
	go build -o ./bin/goblog main.go

modules:
	@go mod tidy

clean:
	@rm -rf bin
