.PHONY:
.SILENT:

build:
	go build -o ../.bin/app .

run: build
	../.bin/app

test:
	go test ./... -coverprofile /dev/null