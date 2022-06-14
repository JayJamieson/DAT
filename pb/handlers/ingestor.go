package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewIngestor() func(ctx context.Context, s3Event events.S3Event) {
	return func(ctx context.Context, s3Event events.S3Event) {
		fmt.Println("Got pricebook, and starting database building")
		for _, record := range s3Event.Records {
			s3 := record.S3
			// TODO process csv into sqlite database and upload s3 database bucket
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		}
	}
}
