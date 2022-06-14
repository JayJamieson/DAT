package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/mattn/go-sqlite3"
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
	if os.Getenv("PB_ENV") == "production" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

type Downloader interface {
	Download(key string, ctx context.Context) error
}

type NoOpDownloader struct {
}

func (noop *NoOpDownloader) Download(key string, ctx context.Context) error {
	return nil
}

func NewSearcher(connectionString string, downloader Downloader) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

		err := downloader.Download("pb.db", ctx)

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		// entries, err := os.ReadDir("/tmp")

		// if err != nil {
		// 	log.Fatalf("error reading directory %v", err)
		// 	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		// }
		// log.Println(len(entries))

		// for _, item := range entries {
		// 	info, _ := item.Info()
		// 	log.Printf("%v, %v, %v", item.Name(), info.Mode(), info.Size())
		// }

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

func getResultLimit(queryStringParameters map[string]string) (int, error) {
	val, ok := queryStringParameters["limit"]

	if !ok {
		return 10, nil
	}
	limit, err := strconv.Atoi(val)

	return limit, err
}
