#!/bin/bash

case $1 in
    -h | --help ) echo "usage: $(basename $0) start|stop|connect"; exit;;
esac


if [ $# -ne 1 ]; then
    2>&1 echo "error: bad number of arguments"
    exit 1
fi

name=gosec-pg

case $1 in
    start)
	docker run \
	    -d \
	    -e POSTGRES_PASSWORD=s3cr3t \
	    -p 5432:5432 \
	    --name gosec-pg \
	    -v $PWD/sql/schema:/docker-entrypoint-initdb.d \
	    postgres:17-alpine;;
    stop) docker rm -f ${name};;
    connect) docker exec -it ${name} psql -U postgres
esac
