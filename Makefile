
run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o kc

test:
	@go test -v