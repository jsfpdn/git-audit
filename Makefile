BIN=./bin/git-audit
FLAGS=-race

build:
	CGO_ENABLED=0 go build -o $(BIN)

install:
	go install $(FLAGS)

fmt:
	gofmt -l .

lint: fmt
	golangci-lint run

test:
	go test $(FLAGS) -v ./... -count=1

docker:
	docker build --tag josefpodanyml/git-audit .

docker-push: docker
	docker push josefpodanyml/git-audit:latest

clean:
	rm $(BIN)

gen-proto:
	protoc -I . ./proto/changelog.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=./proto
