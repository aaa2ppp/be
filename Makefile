all: vet lint test

lint:
	@golangci-lint run --print-issued-lines=false --out-format=colored-line-number ./...
	@echo "✓ lint"

vet:
	@go vet ./...
	@echo "✓ vet"

test:
	@go test ./...
	@echo "✓ test"
