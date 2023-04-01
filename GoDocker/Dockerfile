FROM golang:1.20-alpine

WORKDIR /go/src/goapp
COPY . .

RUN go mod tidy
RUN go build -o /goapp

CMD [“/goapp”]