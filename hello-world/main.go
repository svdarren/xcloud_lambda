package main

import (
	"fmt"
	"hello-world/api"
	"net/http"
)

func handle(r *api.IncomingRequest) (status int, body string, err error) {
	var greeting string

	name := r.Cloud
	if name == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", name)
	}

	return http.StatusOK, greeting, nil
}

func main() {
	api.CloudSpecificSetup(handle)
}
