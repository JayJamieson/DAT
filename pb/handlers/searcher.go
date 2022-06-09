package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewSearcher() func(ctx context.Context, s3Event events.S3Event) {
	return func(ctx context.Context, s3Event events.S3Event) {
		fmt.Println("Got Search request")
	}
}
