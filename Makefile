all: clean build test run install goreleaser-release goreleaser-dry-run goreleaser-dry-run

REPO_NAME = github.com/hoto/template-go-web-server

clean:
	go clean
	rm -rf bin/ dist/

dependencies:
	go mod download
	go mod tidy
	go mod verify

build: dependencies
	go build -ldflags="-X '${REPO_NAME}/config.Version=0.0.0' -X '${REPO_NAME}/config.ShortCommit=HASH' -X '${REPO_NAME}/config.BuildDate=DATE'" -o bin/template-go-web-server

test:
	go test -v ./...

run: clean build
	./bin/template-go-web-server $(arg)

install: clean build
	go install -v ./...

goreleaser-release: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash

goreleaser-dry-run: clean dependencies
	curl -sL https://git.io/goreleaser | VERSION=v0.137.0 bash -s -- --skip-publish --snapshot --rm-dist

goreleaser-dry-run-local: dependencies
	goreleaser release --skip-publish --snapshot --rm-dist

