FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY services/go/go.mod ./go.mod
COPY services/go/internal ./internal
COPY services/go/common ./common
COPY services/go/order-service ./order-service

RUN go build -o /order-service ./order-service

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /order-service /usr/local/bin/order-service
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["order-service"]


