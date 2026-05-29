
.PHONY: run deps build test release deploy

run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o kc

test:
	@go test -v

release: deploy

deploy:
	@command -v goreleaser >/dev/null 2>&1 || { echo "❌ Error: goreleaser is not installed."; exit 1; }
	@command -v gh >/dev/null 2>&1 || { echo "❌ Error: GitHub CLI (gh) is not installed."; exit 1; }
	@gh auth token >/dev/null 2>&1 || { echo "❌ Error: gh is not authenticated. Run: gh auth login --scopes repo"; exit 1; }
	@KC_VERSION=$$(sed -nE 's/^const KcVersion = "(.*)"/\1/p' src/version.go); \
	if [ -z "$$KC_VERSION" ]; then echo "Error: Could not extract version from src/version.go"; exit 1; fi; \
	echo "Releasing version $$KC_VERSION 🚀..."; \
	if git rev-parse "$$KC_VERSION" >/dev/null 2>&1; then \
		echo "Tag $$KC_VERSION already exists locally."; \
	else \
		git tag -a $$KC_VERSION -m "Release $$KC_VERSION"; \
	fi; \
	git push origin $$KC_VERSION; \
	GITHUB_TOKEN=$$(gh auth token) goreleaser release
