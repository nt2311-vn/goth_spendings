BINARY_NAME=spendings
ifeq ($(OS),Windows_NT)
	BINARY_NAME=spendings.exe
endif

.PHONY: build


build:
	go mod tidy && \
		templ generate && \
	bunx tailwindcss build -i static/css/style.css -o static/css/tailwind.css && \
	go build -o ./bin/${BINARY_NAME} ./cmd/main.go && \
	air -c .air.toml

clean:
	go clean
