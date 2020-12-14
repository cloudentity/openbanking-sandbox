FROM golang:1.15

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY client client
COPY models models

RUN go mod download

COPY bank .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
