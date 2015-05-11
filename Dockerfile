FROM golang:latest

RUN go get github.com/codegangsta/gin

EXPOSE 3000