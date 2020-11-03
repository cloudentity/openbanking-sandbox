FROM golang:1.15-alpine

WORKDIR /app

COPY tpp .
COPY vendor vendor
COPY go.mod .
COPY go.sum .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
