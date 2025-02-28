FROM golang:1.24-alpine

COPY . .

RUN go mod download

RUN go build -o main .

VOLUME /config

CMD ["./main"]