#!/bin/bash -xe

docker run --rm --name pipeline-receiver \
    --network host \
    -e RABBITMQ_URL="amqp://guest:guest@141.144.239.9:5672" \
    pipeline-receiver
