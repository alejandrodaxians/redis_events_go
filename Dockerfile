FROM golang:1.21.5-bullseye AS build

WORKDIR /go/src/rediskeyspace

COPY db ./db
COPY main.go ./
COPY go.mod ./
COPY go.sum ./

RUN ls -la

RUN go mod download

RUN go build -o app .
RUN ls -la

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /go/src/mssMicroservice/logs /logs
COPY --from=build /go/src/mssMicroservice/app /app
CMD ["/app"]