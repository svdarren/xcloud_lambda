//go:build !aws

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type httpContext struct {
	// When the handler needs other data, add it here
	cloud string
}

func (c *httpContext) httpHandler(w http.ResponseWriter, req *http.Request) {
	r := incomingRequest{
		req,
		c.cloud,
	}
	status, body, err := r.handle()
	_ = status
	_ = err
	fmt.Fprint(w, body)
}

func cloudSpecificSetup() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	ctx := httpContext{
		cloud: os.Getenv("CLOUD_PROVIDER"),
	}
	log.Printf("CLOUD_PROVIDER set to %s", ctx.cloud)
	http.HandleFunc("/", ctx.httpHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
