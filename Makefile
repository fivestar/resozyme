.PHONY: test
test:
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: run-example
run-example:
	@go run _examples/main.go
