# Go and Docker

## Starting docker on a mac
```
boot2docker start
```

## Build image from our dockerfile
```
docker build -t answer_to_life .
```

## Create redis container
```
docker run -d -it --name atol-redis -p 6380:6379 redis:latest
```

## Create container from our custom docker image and link it to redis
```
docker run -d -it --name answer-to-life --link atol-redis:redis -p 3000:3000 -v $(pwd):/var/www answer_to_life
```

## Start app
```
docker exec -it answer-to-life bash -c "gin -a 4200"
```