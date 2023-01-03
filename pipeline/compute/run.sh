#!/bin/bash -xe

docker run --rm --name pipeline-compute \
    --network host \
    -e KAFKA_URL=127.0.0.1:29092 \
    -e RABBITMQ_URL="amqp://guest:guest@127.0.0.1:5672" \
    pipeline-compute
