FROM golang:latest

RUN go get github.com/codegangsta/gin

# CMD "(cd /var/www/ && gin -a 4200)"

EXPOSE 3000