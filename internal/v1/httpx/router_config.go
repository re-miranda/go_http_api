package httpx

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

type	Router_config []struct {
	name string
	router *httprouter.Router
}

func	(cfg *Router_config)LoadFromFile(file_path string) {
	// placeholder test of two servers
	if file_path == "Default" {
		cfg.newRouterConfig(file_path)
		cfg.newRouterConfig(file_path)
	}
}

func (router *Router_config)newRouterConfig(params string) {
	router_config := newRouter()

	router_config.GET("/healthz", handlers.HealthzHandler)
	router_config.GET("/v1/ping", handlers.PingHandler)
	router_config.POST("/v1/reverse", handlers.ReverseHandler)

	*router = append(*router, struct{
		name string
		router *httprouter.Router
	}{
		name: params,
		router: router_config,
	})
}

func	newRouter() *httprouter.Router{
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true

	return mux
}
