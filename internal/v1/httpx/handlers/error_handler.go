package handlers

import (
	"net/http"
	"encoding/json"
)

type APIError struct {
    Status   string      `json:"error"`
    Details interface{} `json:"details,omitempty"`
}

func	APIerrorJSON(w http.ResponseWriter, status string, code int, details...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if len(details) != 0 {
		json.NewEncoder(w).Encode(APIError{
			Status:   status,
			Details: details,
		})
		return
	}
	json.NewEncoder(w).Encode(APIError{
		Status:   status,
	})
}
