#!/bin/bash

rm ./bin/main

docker run -it -v $(pwd):/project aws go build -o ./bin/main ./cmd/searcher/main.go

# zip -j function.zip ./bin/main ./bin/pb.db

rm function.zip

zip -j function.zip ./bin/main

 aws lambda update-function-code --function-name pb-query --zip-file fileb://function.zip