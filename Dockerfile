FROM golang:1.16

WORKDIR /go/src/app
COPY . ./src
WORKDIR /go/src/app/src/src
CMD make run
