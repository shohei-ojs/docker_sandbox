FROM golang:latest

WORKDIR /go/src/app

EXPOSE 5000

ENTRYPOINT [ "go", "run", "main.go" ]