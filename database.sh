#!/bin/bash

# FROM mysql:8.0
#
# LABEL maintainer="Debezium Community"
#
# COPY mysql.cnf /etc/mysql/conf.d/
# COPY inventory.sql /f/
echo "Starting database systems"

echo "Starting redis"
docker run -p 8080:6379 --name redis -h redis -d redis

case $2 in
    "mysql")
        echo "Starting zookeeper"
        docker run -d --name zookeeper -p 2181:2181 -p 2888:2888 -p 3888:3888 debezium/zookeeper:1.8

        echo "Starting kafka"
        docker run -d --name kafka -p 9092:9092 --link zookeeper:zookeeper debezium/kafka:1.8

        echo "Starting mysql"
        docker run -d --name mysql -h mysql -p 3307:3306 -v $PWD/mysql:/etc/mysql/conf.d -v $PWD/sql:/docker-entrypoint-initdb.d -e MYSQL_ROOT_PASSWORD=debezium -e MYSQL_USER=mysqluser -e MYSQL_PASSWORD=mysqlpw mysql:8.0
        ;;
    "postgres")
        echo "Starting postgress"
        docker run -d --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres debezium/example-postgres
esac
