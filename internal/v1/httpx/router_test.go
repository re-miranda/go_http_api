package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func setupTestRouter() *httprouter.Router {
	mux := httprouter.New()

	// Custom NotFound handler
	mux.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.APIErrorJSON(w, "Not Found", http.StatusNotFound)
	})

	// Custom MethodNotAllowed handler
	mux.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.APIErrorJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	mux.GET("/healthz", handlers.HealthzHandler)
	mux.GET("/v1/ping", handlers.PingHandler)
	mux.POST("/v1/reverse", handlers.ReverseHandler)

	return mux
}

func TestNotFoundHandler(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/non-existent", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["error"] != "Not Found" {
		t.Errorf("Expected error 'Not Found', got %v", response["error"])
	}

	// Check Content-Type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got %s", ct)
	}
}

func TestMethodNotAllowedHandler(t *testing.T) {
	router := setupTestRouter()

	// Test DELETE on /healthz (which only accepts GET)
	req := httptest.NewRequest("DELETE", "/healthz", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rr.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["error"] != "Method not allowed" {
		t.Errorf("Expected error 'Method not allowed', got %v", response["error"])
	}

	// Check Content-Type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got %s", ct)
	}
}

func TestHealthzEndpoint(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["status"] != "ok" {
		t.Errorf("Expected status 'ok', got %v", response["status"])
	}
}

func TestPingEndpoint(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest("GET", "/v1/ping", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	if rr.Body.String() != "pong" {
		t.Errorf("Expected body 'pong', got %s", rr.Body.String())
	}
}