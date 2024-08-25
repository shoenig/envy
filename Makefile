NAME = envy

default: build

.PHONY: build
build: clean
	@echo "--> Build ..."
	CGO_ENABLED=0 go build -o output/$(NAME)

.PHONY: clean
clean:
	@echo "--> Clean ..."
	rm -rf dist output/$(NAME)

.PHONY: test
test:
	@echo "--> Test Go Sources ..."
	go test -race ./...

.PHONY: vet
vet:
	@echo "--> Vet Go Sources ..."
	go vet ./...

.PHONY: copywrite
copywrite:
	@echo "--> Checking Copywrite ..."
	copywrite \
		--config .github/workflows/scripts/copywrite.hcl headers \
		--spdx "MIT"

.PHONY: lint
lint: vet
	@echo "--> Lint ..."
	@golangci-lint run --config .github/workflows/scripts/golangci.yaml

.PHONY: release
release:
	@echo "--> RELEASE ..."
	envy exec gh-release goreleaser release --clean
	$(MAKE) clean
