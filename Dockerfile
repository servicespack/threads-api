FROM golang:1.21.4

WORKDIR /go/src/app

COPY . .

RUN go build -o app
RUN chmod +x app

EXPOSE 8080

CMD ["./app"]
