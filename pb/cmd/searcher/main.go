package main

import (
	"github.com/JayJamieson/pb"
	"github.com/JayJamieson/pb/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.NewSearcher("/tmp/pricebook.db", &pb.S3Downloader{}))
}
