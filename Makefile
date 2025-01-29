
run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o kc

test:
	@go test -v

release:
	@command -v goreleaser >/dev/null 2>&1 || { echo "❌ Error: goreleaser is not installed."; exit 1; }
	@KC_VERSION=$$(sed -nE 's/^const KcVersion = "(.*)"/\1/p' src/version.go); \
	if [ -z "$$KC_VERSION" ]; then echo "Error: Could not extract version from src/version.go"; exit 1; fi; \
	echo "Releasing version $$KC_VERSION 🚀..."; \
	git tag -a $$KC_VERSION -m "Release $$KC_VERSION"; \
	git push origin $$KC_VERSION; \
	goreleaser release
