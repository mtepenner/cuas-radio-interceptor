FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN go build -o /signal-classifier ./cmd/classifier

FROM alpine:3.20
WORKDIR /app
COPY --from=build /signal-classifier /usr/local/bin/signal-classifier

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/signal-classifier"]
