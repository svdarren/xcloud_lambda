//go:build aws

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func (c *httpContext) awsHandler(_ context.Context, req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
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
	r := IncomingRequest{
		httpReq,
		c.cloud,
	}
	resp.StatusCode, resp.Body, err = c.clientHandler(&r)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		return resp, err
	}
	return resp, nil
}

func CloudSpecificSetup(handler RequestHandler) {
	c := httpContext{
		cloud:         os.Getenv("CLOUD_PROVIDER"),
		clientHandler: handler,
	}
	log.Printf("CLOUD_PROVIDER set to %s", c.cloud)
	lambda.StartHandlerFunc(c.awsHandler)
}
