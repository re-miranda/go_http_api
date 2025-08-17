package httpx

import (
	"fmt"
	"net/http"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func Router(config string) error{
	fmt.Println(config)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	mux.HandleFunc("/v1/ping", handlers.PingHandler)
	mux.HandleFunc("/v1/reverse", handlers.ReverseHandler)

	srv := http.Server {
		Addr: ":8080",
		Handler: mux,
	}
	fmt.Println("Server is starting")
	return srv.ListenAndServe()
}