FROM golang:1.20-buster

RUN go version
ENV $GOPATH=/

WORKDIR /app

COPY ./ /app

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod -x wait-for-postgres.sh

RUN go mod download
RUN go build -o movieapi ./cmd/main.go

CMD ["./movieAPI"]