FROM golang:1.20.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go test ./...
RUN go build -o main

FROM alpine:3.17
WORKDIR /app
# Copy binary go file from builder
COPY --from=builder /app/main /app/
# Copy .env file from builder
COPY --from=builder /app/.env /app/
# CMD main
CMD ["./main"]
# Need env file configuration 