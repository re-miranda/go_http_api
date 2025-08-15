package main

import (
	"fmt"
	"net/http"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func main(){
	http.HandleFunc("/healthz", handlers.HealthzHandler)
	http.HandleFunc("/v1/ping", handlers.PingHandler)
	http.HandleFunc("/v1/reverse", handlers.ReverseHandler)

	fmt.Println("server is running")

	http.ListenAndServe(":8080", nil)
}
