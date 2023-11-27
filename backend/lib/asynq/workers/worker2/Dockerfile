# Build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8081
CMD ["/app/main"]

# RUN stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["/app/main"]
