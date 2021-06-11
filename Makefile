build:
	go build -race -o ./bin/git-audit ./cmd/git-audit/main.go

run:
	go run -race ./cmd/git-audit/main.go

install:
	go install -race ./cmd/***

fmt:
	gofmt -l .

lint:
	golangci-lint run

test:
	go test -race -v ./... -count=1

build-docker:
	docker build --tag git-audit .