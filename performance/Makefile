.PHONY: charts http redis

DOCKER_REDIS=docker run -d --rm --name=redis --env=REDIS_PASSWORD=redispw --volume=/data --workdir=/data -p 55000:6379 redis
DOCKER_STOP=docker stop redis

charts:
	$(DOCKER_REDIS)
	go run main.go; $(DOCKER_STOP)

test: http redis

http:
	go test ./http -count=1

redis:
	$(DOCKER_REDIS)
	go test ./redis -count=1; $(DOCKER_STOP)
