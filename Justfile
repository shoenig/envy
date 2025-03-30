set shell := ["bash", "-u", "-c"]

export scripts := ".github/workflows/scripts"
export GOBIN := `echo $PWD/.bin`

# show available commands
[private]
default:
    @just --list

# tidy up Go modules
[group('build')]
tidy:
    go mod tidy

# run tests across source tree
[group('build')]
tests:
    go test -v -race -count=1 ./...

# ensure copywrite headers present on source files
[group('lint')]
copywrite:
    copywrite \
        --config {{scripts}}/copywrite.hcl headers \
        --spdx "MIT"

# apply go vet command on source tree
[group('lint')]
vet:
    go vet ./...

# apply golangci-lint linters on source tree
[group('lint')]
lint: vet
    $GOBIN/golangci-lint run --config {{scripts}}/golangci.yaml

# locally install build dependencies
[group('build')]
init:
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

# show host system information
[group('build')]
@sysinfo:
    echo "{{os()/arch()}} {{num_cpus()}}c"

