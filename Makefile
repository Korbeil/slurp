# Go parameters
GOCMD=go
BUILD_DIR=build
BINARY_NAME=slurp
BINARY_UNIX=$(BINARY_NAME)_unix
.PHONY: build

all: test build

# Build tasks
build:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)
	$(GOCMD) build -v
	mv $(BINARY_NAME) $(BUILD_DIR)/
build-linux:
	rm -f $(BUILD_DIR)/$(BINARY_UNIX)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -v
	mv $(BINARY_NAME) $(BUILD_DIR)/$(BINARY_UNIX)
clean:
	$(GOCMD) clean
	rm -f $(BUILD_DIR)/$(BINARY_NAME)
	rm -f $(BUILD_DIR)/$(BINARY_UNIX)
install:
	$(GOCMD) install

test:
	$(GOCMD) test -v ./...
deps:
	$(GOCMD) get github.com/Korbeil/slurp/commands
	$(GOCMD) get github.com/Korbeil/slurp/utils/directory
	$(GOCMD) get github.com/Korbeil/slurp/utils/json
	$(GOCMD) get github.com/urfave/cli