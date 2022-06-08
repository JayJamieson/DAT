package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)


func NewIngestor() func(ctx context.Context, s3Event events.S3Event) {
	return func(ctx context.Context, s3Event events.S3Event) {
		fmt.Println("Got email, and starting attachment extraction")
		for _, record := range s3Event.Records {
			s3 := record.S3
			// TODO log event key as structured JSON
			// fetch s3 stream for file
			// TODO extract email header and copy to processing directory using email
			// TODO log email header as structured json
			// TODO log s3 attachment to dynamo
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		}
	}
}
