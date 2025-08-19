package httpx

import (
	"fmt"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func Router(config string, done chan string) error{
	fmt.Println(config)

	// Create and set multiplexer (router)
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true
	mux.GET("/healthz", handlers.HealthzHandler)
	mux.GET("/v1/ping", handlers.PingHandler)
	mux.POST("/v1/reverse", handlers.ReverseHandler)

	// Create server
	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	done <- "Server is starting"
	return srv.ListenAndServe()
}