#!/bin/bash -xe

docker run --rm --name pipeline-sensor \
    --network host \
    -e KAFKA_URL=127.0.0.1:29092 \
    pipeline-sensor
