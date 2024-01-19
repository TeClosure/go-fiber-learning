FROM golang:1.21.6-alpine3.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
