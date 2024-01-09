# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]
