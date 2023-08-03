package main

import (
	"fmt"
	"net/http"
)

type incomingRequest struct {
	request *http.Request
	cloud   string
}

func (r incomingRequest) handle() (status int, body string, err error) {
	var greeting string

	name := r.cloud
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
