#Building State
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum  ./
COPY . .
RUN go mod download
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]