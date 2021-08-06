FROM golang:1.16-alpine

WORKDIR /sherlock

COPY . ./

RUN go mod download

RUN go build -o sherlock

ENTRYPOINT ["./sherlock"]

