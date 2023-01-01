.PHONY: start

DOCKER_KAFKA=cd ../compose/kafka && docker compose up -d
DOCKER_STOP=docker ps -aq | xargs docker stop | xargs docker rm

start:
	$(DOCKER_KAFKA)
	sleep 10
	( \
		cd sensor; \
		source venv/bin/activate; \
		pip install -r requirements.txt; \
		python src/main.py; \
	);
	$(DOCKER_STOP)
