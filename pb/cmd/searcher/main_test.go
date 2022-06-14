package main_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/JayJamieson/pb/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/matryer/is"
)

var pbPath = "/tmp/pb.db"

type pricebookDownloader struct{}

func (mock pricebookDownloader) Download(key string, ctx context.Context) error {
	return nil
}

func TestSearcherLimitOne(t *testing.T) {
	_is := is.New(t)

	d := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "495b12a8-xmpl-4eca-8168-160484189f99",
		InvokedFunctionArn: "arn:aws:lambda:ap-southeast-2:123456789012:function:pb-searcher",
	})

	var event events.APIGatewayProxyRequest

	event.QueryStringParameters = map[string]string{
		"q":     "conduit",
		"limit": "1",
	}

	handler := handlers.NewSearcher(pbPath, &pricebookDownloader{})

	response, err := handler(ctx, event)

	_is.NoErr(err)

	var results handlers.SearchResults
	err = json.Unmarshal([]byte(response.Body), &results)
	_is.NoErr(err)

	_is.Equal(len(results), 1)
}

func TestSearcherUsesDefaultLimit(t *testing.T) {
	_is := is.New(t)

	d := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "495b12a8-xmpl-4eca-8168-160484189f99",
		InvokedFunctionArn: "arn:aws:lambda:ap-southeast-2:123456789012:function:pb-searcher",
	})

	var event events.APIGatewayProxyRequest

	event.QueryStringParameters = map[string]string{
		"q": "conduit",
	}

	handler := handlers.NewSearcher(pbPath, &pricebookDownloader{})

	response, err := handler(ctx, event)

	_is.NoErr(err)

	var results handlers.SearchResults
	err = json.Unmarshal([]byte(response.Body), &results)
	_is.NoErr(err)

	_is.Equal(len(results), 10)
}

func TestSearcherSearchResultCorrect(t *testing.T) {
	type test struct {
		value   string
		isValid bool
	}

	cases := []test{
		{value: "conduit", isValid: true},
		{value: "tape", isValid: true},
	}

	_is := is.New(t)
	d := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "495b12a8-xmpl-4eca-8168-160484189f99",
		InvokedFunctionArn: "arn:aws:lambda:ap-southeast-2:123456789012:function:pb-searcher",
	})

	var event events.APIGatewayProxyRequest

	handler := handlers.NewSearcher(pbPath, &pricebookDownloader{})

	for _, query := range cases {
		event.QueryStringParameters = map[string]string{
			"q": query.value,
		}

		response, err := handler(ctx, event)
		_is.NoErr(err)

		var results handlers.SearchResults

		err = json.Unmarshal([]byte(response.Body), &results)

		_is.NoErr(err)
		_is.True(strings.Contains(strings.ToLower(results[0].ProductName), query.value))
	}
}
