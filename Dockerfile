FROM golang:latest

RUN go get github.com/codegangsta/gin
RUN go get menteslibres.net/gosexy/redis

WORKDIR /var/www

CMD "(cd /var/www && gin -a 4200)"

EXPOSE 3000