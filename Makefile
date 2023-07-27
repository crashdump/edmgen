.PHONY: build clean test
.DEFAULT_GOAL=test

build: test
	go build -v -o bin/edmgen/edmgen cmd/edmgen/*.go

clean:
	rm -f bin/edmgen

test:
	go mod tidy
	staticcheck ./...
	go test -v ./...

integration-test:
	go test -v ./... -tags=integration

release: test integration-test build
	go test all -v