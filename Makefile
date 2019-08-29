GO := GO111MODULE=on go

.PHONY: test
test:
	@$(GO) test ./...

.PHONY: lint
lint:
	@golint ./...
	@$(GO) vet ./...
