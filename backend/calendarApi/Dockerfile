FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./api/cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 9403
CMD ["/app/main"]