#!/bin/bash -xe

docker run --rm --name pipeline-compute \
    --network host \
    pipeline-compute
