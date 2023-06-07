FROM golang:alpine3.17

WORKDIR /app

COPY . .

RUN go mod download

RUN go test ./...

RUN go build -o contoh

CMD ["./contoh"]

# Neen env file configuration 