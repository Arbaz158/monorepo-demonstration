FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY services/go/go.mod ./go.mod
COPY services/go/internal ./internal
COPY services/go/common ./common
COPY services/go/payment-service ./payment-service

RUN go build -o /payment-service ./payment-service

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /payment-service /usr/local/bin/payment-service
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["payment-service"]


