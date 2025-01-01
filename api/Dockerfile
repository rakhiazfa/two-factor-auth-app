# Stage 1: Build the Go application
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/api cmd/api/main.go

# Stage 2: Create a minimal image with Alpine and the compiled Go binary
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/api ./bin/api
COPY --from=builder /app/config.json ./

ENV GIN_MODE=release

EXPOSE 8080
CMD ["./bin/api"]