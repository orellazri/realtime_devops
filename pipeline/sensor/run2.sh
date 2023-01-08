#!/bin/bash -xe

docker run --rm --name pipeline-sensor \
    --network host \
    -e KAFKA_URL=144.24.182.241:29092 \
    pipeline-sensor
