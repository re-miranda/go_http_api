package httpx

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func newRouter(params string) *httprouter.Router{
	router := newMux()

	router.GET("/healthz", handlers.HealthzHandler)
	router.GET("/v1/ping", handlers.PingHandler)
	router.POST("/v1/reverse", handlers.ReverseHandler)

	return router
}

func	newMux() *httprouter.Router{
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true

	return mux
}
