package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/mattn/go-sqlite3"
)

func NewSearcher() func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

		for key, value := range request.QueryStringParameters {
			log.Printf("    %s: %s\n", key, value)
		}

		db, err := sql.Open("sqlite3", "pb.db")

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query("select rowid,* from jarussell_fts where jarussell_fts match ? limit 10", request.QueryStringParameters["q"])

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		response := make(map[string]string)

		for rows.Next() {
			var id string
			var product_code string
			var product_name string
			var unit_type string
			var cost_price string
			var retail_price string
			var trade_price string
			var search_values string
			var supplier_sk string

			err = rows.Scan(&id, &product_code, &product_name, &unit_type, &cost_price, &retail_price, &trade_price, &search_values, &supplier_sk)

			if err != nil {
				log.Fatal(err)
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
			}
			response[id] = fmt.Sprintf(
				"%v, %v, %v, %v, %v, %v, %v, %v, %v",
				id, product_code, product_name, unit_type, cost_price, retail_price, trade_price, search_values, supplier_sk)
			fmt.Println(id, product_name)
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
