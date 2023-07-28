//go:build azure

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func azureHandler(w http.ResponseWriter, r *http.Request) {
	status, body, err := handler(r)
	_ = status
	_ = err
	fmt.Fprint(w, body)
}

func cloudSpecificSetup() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/", azureHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
