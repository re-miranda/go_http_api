package httpx

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func Router(routes string) *httprouter.Router{
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true

	mux.GET("/healthz", handlers.HealthzHandler)
	mux.GET("/v1/ping", handlers.PingHandler)
	mux.POST("/v1/reverse", handlers.ReverseHandler)

	return mux
}
