test:
	docker run -d --rm --name=redis --env=REDIS_PASSWORD=redispw --volume=/data --workdir=/data -p 55000:6379 redis
	go test ./... -count=1
	docker stop redis
