FROM golang:1.21.5-bullseye AS build

WORKDIR /rediskeyspace

COPY db ./db
COPY *.go ./
COPY go.mod go.sum ./

RUN ls -la

RUN go mod download

RUN go build -o rediskeyspace .
RUN ls -la

CMD ["/rediskeyspace"]