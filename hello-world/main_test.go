package main

import (
	"hello-world/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		name           string
		request        api.IncomingRequest
		handler        api.RequestHandler
		expectedStatus int
		expectedBody   string
		expectedError  error
	}{
		{
			name: "Basic",
			request: api.IncomingRequest{
				Request: httptest.NewRequest(http.MethodGet, "/", nil),
				Cloud:   "",
			},
			handler:        handle,
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, world!\n",
			expectedError:  nil,
		},
		{
			name: "Include Cloud Provider variable",
			request: api.IncomingRequest{
				Request: httptest.NewRequest(http.MethodGet, "/", nil),
				Cloud:   "Unit Test",
			},
			handler:        handle,
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Unit Test!\n",
			expectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			status, response, err := testCase.handler(&testCase.request)
			if err != testCase.expectedError {
				t.Errorf("Expected error %v, but got %v", testCase.expectedError, err)
			}

			if status != testCase.expectedStatus {
				t.Errorf("Expected status code %v, but got %v", testCase.expectedStatus, status)
			}

			if response != testCase.expectedBody {
				t.Errorf("Expected response %v, but got %v", testCase.expectedBody, response)
			}
		})
	}
}
