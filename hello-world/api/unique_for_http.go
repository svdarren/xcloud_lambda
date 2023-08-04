//go:build !aws

package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (c *httpContext) httpHandler(w http.ResponseWriter, req *http.Request) {
	r := IncomingRequest{
		req,
		c.cloud,
	}
	status, body, err := c.clientHandler(&r)
	_ = status
	_ = err
	fmt.Fprint(w, body)
}

func CloudSpecificSetup(handler RequestHandler) {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	ctx := httpContext{
		cloud:         os.Getenv("CLOUD_PROVIDER"),
		clientHandler: handler,
	}
	log.Printf("CLOUD_PROVIDER set to %s", ctx.cloud)
	http.HandleFunc("/", ctx.httpHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
