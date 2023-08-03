package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		name           string
		request        incomingRequest
		expectedStatus int
		expectedBody   string
		expectedError  error
	}{
		{
			name: "Basic",
			request: incomingRequest{
				request: httptest.NewRequest(http.MethodGet, "/", nil),
				cloud:   "",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, world!\n",
			expectedError:  nil,
		},
		{
			name: "Include Cloud Provider variable",
			request: incomingRequest{
				request: httptest.NewRequest(http.MethodGet, "/", nil),
				cloud:   "Unit Test",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Unit Test!\n",
			expectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			status, response, err := testCase.request.handle()
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
