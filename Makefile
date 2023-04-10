.PHONY: lint run


lint:
	golangci-lint run --timeout 5m

run:
	go run cmd/main.go

