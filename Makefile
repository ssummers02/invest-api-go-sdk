
all: format lint-fix

format:
	gofmt -s -w . && \
	go vet ./... && \
	go mod tidy

gen-proto:
	protoc -I=./contracts --go_out=plugins=grpc:pkg/investapi contracts/*.proto


lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0

# [CI\CD] lint code
lint:
	golangci-lint run
# lint and auto-fix possible problems
lint-fix:
	golangci-lint run --fix

