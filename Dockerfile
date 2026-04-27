FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# install air (optional for dev)
RUN go install github.com/air-verse/air@latest

EXPOSE 8082

CMD ["air"]