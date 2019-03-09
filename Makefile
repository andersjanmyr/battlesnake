
.PHONY: install
install:
	go get -t ./...

build:
	go build

.PHONY: run
run: build
	./battlesnake server

.PHONY: watch
watch:
	./scripts/restart.sh
	fswatch -e ".*" -i "\\.go$$" -0 . | xargs -0 -n 1 ./scripts/restart.sh

test:
	go test ./...
.PHONY: test

fmt:
	@echo ">> Running Gofmt.."
	gofmt -l -s -w .
