#!/bin/bash

docker stop redis && docker rm redis
docker stop mysql && docker rm mysql
docker stop postgres && docker rm postgres
docker stop kafka && docker rm kafka
docker stop zookeeper && docker rm zookeeper
docker stop debezium && docker rm debezium
rm ./data/offsets.dat
