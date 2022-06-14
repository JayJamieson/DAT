package pb

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	priceBookBucketName = "price-books"
	dbPath              = "/tmp/pricebook.db"
)

type S3Downloader struct {
}

func (downloader *S3Downloader) Download(path string, ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(priceBookBucketName),
		Key:    aws.String(path),
	}

	object, err := client.GetObject(ctx, getObjectInput)

	if err != nil {
		log.Fatalf("failed to download pricebook db, %v", err)
		return err
	}

	data, err := ioutil.ReadAll(object.Body)
	if err != nil {
		log.Fatalf("failed to download s3 body, %v", err)
		return err
	}

	err = ioutil.WriteFile(dbPath, data, 0666)

	if err != nil {
		log.Fatalf("failed to download/write pricebook db, %v", err)
		return err
	}

	return nil
}
