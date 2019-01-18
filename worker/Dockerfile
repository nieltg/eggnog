FROM golang:1.11-stretch AS build

WORKDIR /go/src/github.com/nieltg/eggnog/worker
COPY ./ ./

RUN go get ./... \
  && go install -v ./...

FROM ubuntu:18.04 AS final

WORKDIR /opt/worker
COPY --from=build /go/bin /opt/worker/bin

CMD bin/worker
