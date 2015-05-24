FROM golang:latest

RUN go get github.com/codegangsta/gin
RUN go get menteslibres.net/gosexy/redis

WORKDIR /var/www

EXPOSE 3000