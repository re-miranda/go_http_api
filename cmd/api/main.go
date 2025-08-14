package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/re-miranda/go_http_api/internal/v1/reverse"
)

func main(){
	healthz := func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}
	http.HandleFunc("/healthz", healthz)

	ping := func (w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "pong")
	}
	http.HandleFunc("/v1/ping", ping)

	http.HandleFunc("/v1/reverse", reverse.ReverseHandler)

	fmt.Println("server is running")

	http.ListenAndServe(":8080", nil)
}
