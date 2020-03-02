.PHONY: build test

build:
	go build -i -o artifacts/compete github.com/LukeJelly/Cell-Compete-Go

test:
	go list ./... | grep -v vendor | xargs go test -race
