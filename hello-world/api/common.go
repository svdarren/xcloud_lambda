package api

import "net/http"

type IncomingRequest struct {
	Request *http.Request
	Cloud   string
}

type RequestHandler func(*IncomingRequest) (status int, body string, err error)

type httpContext struct {
	// When the handler needs other data, add it here
	cloud         string
	clientHandler RequestHandler
}
