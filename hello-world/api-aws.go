//go:build aws

package main

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

type awsContext struct {
	// When the handler needs other data, add it here
	cloud string
}

func (a *awsContext) awsHandler(_ context.Context, req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
	log.Printf("%d", a.cloud)
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
	r := incomingRequest{
		httpReq,
		a.cloud,
	}
	resp.StatusCode, resp.Body, err = r.handle()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		return resp, err
	}
	return resp, nil
}

func cloudSpecificSetup() {
	a := awsContext{
		cloud: os.Getenv("CLOUD_PROVIDER"),
	}
	lambda.StartHandlerFunc(a.awsHandler)
}
