build-docker: 
	CGO_ENABLED=0 GOOS=linux go build -o bin/docker-golang
run:
	./bin/docker-golang

build: 
	@go build -o bin/docker-golang

dev: build
	@./bin/docker-golang

test: 
	@go test -v ./...
