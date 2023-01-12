.PHONY: charts http redis kafka

DOCKER_REDIS=cd ../compose/redis && docker compose up -d
DOCKER_KAFKA=cd ../compose/kafka && docker compose up -d
DOCKER_STOP=docker ps -aq | xargs docker stop | xargs docker rm

charts:
	$(DOCKER_REDIS)
	$(DOCKER_KAFKA)
	sleep 10
	go run main.go

test: http redis kafka

http:
	go test ./http -count=1

redis:
	$(DOCKER_REDIS)
	go test ./redis -count=1

kafka:
	$(DOCKER_KAFKA)
	sleep 10
	go test ./kafka -count=1
