#!/bin/bash

dirs=($(find . -mindepth 1 -type d))
for dir in ${dirs[@]}; do
	pushd $dir
	docker compose up -d
	popd
done
