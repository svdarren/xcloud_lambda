package main

import (
	"fmt"
	"net/http"
)

type Query struct {
	Name string `json:"name"`
}

type Request struct {
	Body    string                 `json:"body"`
	Headers map[string]interface{} `json:"headers"`
	Query   Query                  `json:"queryStringParameters"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int
}

func handler(req *http.Request) (status int, body string, err error) {
	var greeting string

	name := req.URL.Query().Get("name")
	if name == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", name)
	}

	return http.StatusOK, greeting, nil
}

func main() {
	cloudSpecificSetup()
}
