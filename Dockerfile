FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# install air (optional for dev)
RUN go install github.com/air-verse/air@latest

EXPOSE 3000

CMD ["air"]