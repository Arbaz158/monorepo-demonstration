FROM golang:1.21-alpine AS builder
WORKDIR /app

# Recreate the Go module layout from services/go
COPY services/go/go.mod ./go.mod
COPY services/go/internal ./internal
COPY services/go/common ./common
COPY services/go/user-service ./user-service

RUN go build -o /user-service ./user-service

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /user-service /usr/local/bin/user-service
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["user-service"]


