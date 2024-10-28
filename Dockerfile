FROM golang:1.22.3-alpine

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]