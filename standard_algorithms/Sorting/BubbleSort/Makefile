

all: lint test build

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.33.0

install-richgo:
	go install github.com/kyoh86/richgo@v0.3.11

lint: install-lint
	golangci-lint run --timeout=3m -v

test: install-richgo
	RICHGO_FORCE_COLOR=1 richgo test -v ./... || RICHGO_FORCE_COLOR=1 richgo test -v ./...
	richgo test -json ./... > report.json

run: 
	echo "Running..."
