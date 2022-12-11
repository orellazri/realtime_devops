.PHONY: redis http

redis:
	docker run -d --rm --name=redis --env=REDIS_PASSWORD=redispw --volume=/data --workdir=/data -p 55000:6379 redis
	go test ./redis -count=1; docker stop redis

http:
	go test ./http -count=1
