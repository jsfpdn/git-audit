FROM golang:alpine AS builder

RUN apk update && apk add go make 

WORKDIR /go/src/git-audit
COPY . .

RUN go get -d -v
RUN make build

FROM alpine 

# Certificates are needed in order to call the github API.
RUN apk add ca-certificates
RUN update-ca-certificates 2>/dev/null || true

WORKDIR /
COPY --from=builder /go/src/git-audit/bin/git-audit /usr/bin/git-audit

ENTRYPOINT ["git-audit"]
