FROM golang:latest

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p tmp

RUN go build -o tmp/main ./app

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]