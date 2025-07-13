# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ARG SERVICE

RUN go build -o /bin/service ./cmd/${SERVICE}

# Final minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/service /app/service

CMD ["/app/service"]
