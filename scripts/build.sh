#!/usr/bin/env bash

# Run listener and builder builds prior to docker compose build
dir="${0%/*}"
cd "$dir/../" || exit

for build in "listener" "builder"; do
	"./$build/scripts/build.sh"
done

docker-compose build
