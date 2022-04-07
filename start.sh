#!/bin/bash

START_DB=${1:-"FALSE"}
DB_NAME=${2:-"mysql"}
DB_TYPE=${3:-"mysql"}

if [ "$START_DB" == "TRUE" ] ; then
    ./database.sh $DB_NAME $DB_TYPE $EXISTING_MYSQL
fi

case $DB_TYPE in
    "mysql")
        echo "Running debezium with mysql"
        docker run -it --name debezium -p 9090:8080 -v $PWD/conf:/debezium/conf -v $PWD/data:/debezium/data --link redis --link $DB_NAME --link kafka debezium/server
        ;;
    "postgres")
        echo "Running debezium with postgres"
        docker run -it --name debezium -p 9090:8080 -v $PWD/conf:/debezium/conf -v $PWD/data:/debezium/data --link redis --link $DB_NAME debezium/server
esac
