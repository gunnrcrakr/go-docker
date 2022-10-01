FROM golang:latest

WORKDIR /go/src/github.com/gunnrcrakr/go-docker

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]