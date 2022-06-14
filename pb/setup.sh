#!/bin/bash

if [ -f "pb.db" ]; then
    rm pb.db
fi

if [ -d "./bin" ]; then 
    rm -rf ./bin
fi

mkdir ./bin
cat ./testdata/setup_fts3.sql | sqlite3 pb.db

cp ./pb.db ./bin
cp ./pb.db /tmp/pb.db

docker build -t aws .