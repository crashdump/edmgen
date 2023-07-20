.PHONY: build clean test

build: test
	go build -o bin/edmctl cmd/*.go

clean:
	rm -f bin/edmctl

test:
	go mod tidy
	staticcheck ./...
	go test ./...

integration-test:
	go test ./...  -tags=integration