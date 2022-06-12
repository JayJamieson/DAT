package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/mattn/go-sqlite3"
)

var (
	priceBookBucketName = "price-books"
	dbPath              = "/tmp/pricebook.db"
)

type SearchResult struct {
	Id           string
	ProductCode  string
	ProductName  string
	UnitType     string
	CostPrice    string
	RetailPrice  string
	TradePrice   string
	SearchValues string
	SupplierSk   string
}

type SearchResults []SearchResult

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewSearcher(connectionString string) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

		err := downloadPriceBook("pb.db", ctx)

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		entries, err := os.ReadDir("/tmp")

		if err != nil {
			log.Fatalf("error reading directory %v", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}
		log.Println(len(entries))

		for _, item := range entries {
			info, _ := item.Info()
			log.Printf("%v, %v, %v", item.Name(), info.Mode(), info.Size())
		}

		db, err := sql.Open("sqlite3", connectionString)

		if err != nil {
			log.Fatalf("error opening db %v", err)
		}

		defer db.Close()

		limit, err := getResultLimit(request.QueryStringParameters)

		if err != nil {
			log.Fatalf("error getting query param limit %v", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}

		// tables, err := db.Query("select tbl_name from main.sqlite_master where type = 'table'")

		// log.Printf("error selecting tables %v", err)

		// if err != nil {
		// 	log.Fatal(err)
		// 	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		// }

		// log.Println("available tables")
		// for tables.Next() {
		// 	var table string

		// 	err := tables.Scan(&table)

		// 	if err != nil {
		// 		log.Fatal(err)
		// 		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		// 	}
		// 	log.Println(table)
		// }
		// log.Println("available tables done")

		rows, err := db.Query("select rowid, * from jarussell_fts where jarussell_fts match ? limit ?", request.QueryStringParameters["q"], limit)

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		response := make(SearchResults, 0, limit)

		for rows.Next() {
			item := SearchResult{}

			err = rows.Scan(
				&item.Id,
				&item.ProductCode,
				&item.ProductName,
				&item.UnitType,
				&item.RetailPrice,
				&item.SearchValues,
				&item.SupplierSk,
				&item.TradePrice,
				&item.UnitType,
			)

			if err != nil {
				log.Fatal(err)
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
			}
			response = append(response, item)
		}

		err = rows.Err()

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		body, err := json.Marshal(response)

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
	}
}

func downloadPriceBook(path string, ctx context.Context) error {
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

func getResultLimit(queryStringParameters map[string]string) (int, error) {
	val, ok := queryStringParameters["limit"]

	if !ok {
		return 10, nil
	}
	limit, err := strconv.Atoi(val)

	return limit, err
}
