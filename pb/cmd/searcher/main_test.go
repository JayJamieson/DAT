package main_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/JayJamieson/pb/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func TestSearcher(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "pb-ingestor")
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "495b12a8-xmpl-4eca-8168-160484189f99",
		InvokedFunctionArn: "arn:aws:lambda:ap-southeast-2:123456789012:function:pb-ingestor",
	})

	var event events.APIGatewayProxyRequest

	event.QueryStringParameters = map[string]string{
		"q": "conduit",
	}

	handler := handlers.NewSearcher("/tmp/pb.db")

	response, _ := handler(ctx, event)
	fmt.Println(response.Body)
}
