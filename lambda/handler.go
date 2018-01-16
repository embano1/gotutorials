package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	GOOS         string `json:"goos"`
	GOMAXPROCS   int    `json:"gomaxprocs"`
	NUMCPU       int    `json:"numcpu"`
	NUMGOROUTINE int    `json:"numgoroutine"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	b, err := json.Marshal(Response{
		GOOS:         runtime.GOOS,
		GOMAXPROCS:   runtime.GOMAXPROCS(0),
		NUMCPU:       runtime.NumCPU(),
		NUMGOROUTINE: runtime.NumGoroutine()})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return events.APIGatewayProxyResponse{
			Body:       string(b),
			StatusCode: 200,
			Headers: map[string]string{
				"Runtime": "Go",
			},
		},
		nil
}

func main() {
	lambda.Start(handleRequest)
}
