FROM golang:alpine

WORKDIR /go/src/git-audit
COPY . .

RUN apk update && apk add go make
RUN make install

CMD ["git-audit"]