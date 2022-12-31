#!/bin/bash -xe

docker run --rm --name pipeline-sensor \
    --network host \
    pipeline-sensor
