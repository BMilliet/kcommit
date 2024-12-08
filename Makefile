
run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o kcommit

test:
	@go test -v