FROM golang:1.16-alpine

WORKDIR /sherlock

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o sherlock

ENTRYPOINT ["./sherlock"]

