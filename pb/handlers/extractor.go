package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewExtractor() func(ctx context.Context, s3Event events.S3Event) {
	return func(ctx context.Context, s3Event events.S3Event) {
		fmt.Println("Got raw pricebook, and starting extraction")
		for _, record := range s3Event.Records {
			s3 := record.S3
			//TODO fetch s3 stream for file
			//TODO parse csv to fergus csv format
			//TODO dispatch to ingestion bucket
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		}
	}
}
