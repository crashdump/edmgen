.PHONY: build clean test
.DEFAULT_GOAL=test

build: test
	go build -o bin/edmgen/edmgen cmd/edmgen/*.go

clean:
	rm -f bin/edmgen

test:
	go mod tidy
	staticcheck ./...
	go test ./...

integration-test:
	go test ./...  -tags=integration
