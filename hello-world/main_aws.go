//go:build !azure

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func awsHandler(ctx context.Context, req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		return resp, err
	}
	httpReq, err := http.NewRequest(req.HTTPMethod, req.Path, bytes.NewReader(bodyBytes))
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		return resp, err
	}
	resp.StatusCode, resp.Body, err = handler(httpReq)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		return resp, err
	}
	return resp, nil
}

func main() {
	lambda.Start(awsHandler)
}
