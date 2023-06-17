build: 
	@go build -o bin/docker-golang

run: build
	@./bin/docker-golang

test: 
	@go test -v ./...
