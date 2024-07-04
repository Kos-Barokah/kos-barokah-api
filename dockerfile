FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o app .

EXPOSE 8080

CMD ["/app/app"]