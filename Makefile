BINARY_NAME=spendings
ifeq ($(OS),Windows_NT)
	BINARY_NAME=spendings.exe
endif

.PHONY: build


build:
	go mod tidy && \
		templ generate && \
	go build -o ./bin/${BINARY_NAME} ./cmd/main.go

clean:
	go clean
