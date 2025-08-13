BIN = bin
BIN_NAME = api
PKG = ./cmd/api

build: $(PKG)
	@mkdir -p $(BIN)
	@go build -o $(BIN)/$(BIN_NAME) $(PKG)

run: build
	@$(BIN)/$(BIN_NAME)

all: run

test:
	go test ./internal/morestrings

clean:
	rm -rf $(BIN)

fclean: clean

.PHONY: all build run clean fclean test

