# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=tf-simplify

all: clean build

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

run:
	bin/$(BINARY_NAME)

deps:
	$(GOGET) github.com/spf13/cobra@latest

.PHONY: all build clean run deps
