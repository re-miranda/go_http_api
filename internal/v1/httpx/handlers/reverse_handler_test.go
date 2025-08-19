package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		// 405 Method Not Allowed
		{
			name:           "GET method should return 405",
			method:         "GET",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]interface{}{"error": "Method not allowed"},
		},
		// 400 Bad Request scenarios
		{
			name:           "Empty POST body should return 400",
			method:         "POST",
			body:           "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Bad Request", "details": []interface{}{"EOF"}},
		},
		{
			name:           "Invalid JSON should return 400",
			method:         "POST",
			body:           `{"input":hello}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Bad Request", "details": []interface{}{"invalid character 'h' looking for beginning of value"}},
		},
		{
			name:           "Unknown field should return 400",
			method:         "POST",
			body:           `{"inpput":""}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Bad Request", "details": []interface{}{"json: unknown field \"inpput\""}},
		},
		// 200 OK scenarios
		{
			name:           "Empty input should return empty output",
			method:         "POST",
			body:           `{}`,
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"input": "", "output": ""},
		},
		{
			name:           "Valid input should return reversed output",
			method:         "POST",
			body:           `{"input":"hello"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"input": "hello", "output": "olleh"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/v1/reverse", bytes.NewBufferString(tt.body))
			if tt.method == "POST" {
				req.Header.Set("Content-Type", "application/json")
			}

			rr := httptest.NewRecorder()
			ReverseHandler(rr, req, nil)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for key, expectedValue := range tt.expectedBody {
				if actualValue, ok := response[key]; !ok {
					t.Errorf("Expected key %s not found in response", key)
				} else {
					switch expected := expectedValue.(type) {
					case []interface{}:
						actual, ok := actualValue.([]interface{})
						if !ok {
							t.Errorf("Expected %s to be array, got %T", key, actualValue)
							continue
						}
						if len(actual) != len(expected) {
							t.Errorf("Expected %s length %d, got %d", key, len(expected), len(actual))
							continue
						}
						for i, exp := range expected {
							if actual[i] != exp {
								t.Errorf("Expected %s[%d] to be %v, got %v", key, i, exp, actual[i])
							}
						}
					default:
						if actualValue != expectedValue {
							t.Errorf("Expected %s to be %v, got %v", key, expectedValue, actualValue)
						}
					}
				}
			}
		})
	}
}
