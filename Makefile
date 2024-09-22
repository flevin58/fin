build:
	@go build -o bin/ -ldflags "-s -w"

install:
	@go install -ldflags "-s -w"

all: build
