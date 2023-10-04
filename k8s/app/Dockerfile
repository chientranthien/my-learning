FROM golang:1.17
WORKDIR /go/delivery
COPY main.go .
COPY go.mod .
RUN go build -o /server
CMD ["/server"]
