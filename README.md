# Go and Docker

## Starting docker on a mac
```
boot2docker start
```

## Build image from our dockerfile
```
docker build -t answer_to_life .
```

## Create container from our custom docker image
```
docker run -d -it --name answer-to-life -p 3000:3000 -v $(pwd):/var/www answer_to_life
```

## Start app
```
docker exec -it answer-to-life bash (cd /var/www/ && gin -a 4200)
```
