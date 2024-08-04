.PHONY: build

BINARY_NAME=spendings

build:
	go mod tidy && \
		templ generate && \
	go build -o ./bin/${BINARY_NAME} ./cmd/main.go
