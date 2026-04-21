FROM golang:1.25.4-alpine

WORKDIR /var/www/html
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/go-crm-api

EXPOSE 8080

CMD ["go", "run", "./cmd/go-crm-api/main.go"]