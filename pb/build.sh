#!/bin/bash

docker run -it -v $(pwd):/project aws go build -o ./bin/main ./cmd/searcher/main.go

zip -j function.zip ./bin/main ./bin/pb.db