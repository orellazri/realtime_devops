#!/bin/bash -xe

docker run --rm --name pipeline-receiver \
    --network host \
    -e RABBITMQ_URL="amqp://guest:guest@127.0.0.1:5672" \
    pipeline-receiver
