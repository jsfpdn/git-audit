BIN=./bin/git-audit
FLAGS=-race

build:
	go build $(FLAGS) -o $(BIN)

install:
	go install $(FLAGS)

fmt:
	gofmt -l .

lint:
	golangci-lint run

test:
	go test $(FLAGS) -v ./... -count=1

build-docker:
	docker build --tag git-audit .

clean:
	rm $(BIN)