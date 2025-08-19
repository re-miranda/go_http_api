package httpx

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func Router(config Config) *httprouter.Router{
	fmt.Println("Router config:", config.Config_path)

	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true

	mux.GET("/healthz", handlers.HealthzHandler)
	mux.GET("/v1/ping", handlers.PingHandler)
	mux.POST("/v1/reverse", handlers.ReverseHandler)

	return mux
}
