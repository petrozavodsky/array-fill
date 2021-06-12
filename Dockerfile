FROM golang:latest

WORKDIR /go/src/app
COPY . ./src
WORKDIR /go/src/app/src/src
CMD make run
