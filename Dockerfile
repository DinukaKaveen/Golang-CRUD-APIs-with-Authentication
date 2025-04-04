FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o goapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/goapp .

CMD ["./cmd/main.go"]

# Builds docker image
# docker build -t goapp .

# Run docker image to create the container
# docker run -p 8080:8080 goapp
