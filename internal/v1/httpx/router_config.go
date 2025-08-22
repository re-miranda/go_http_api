package httpx

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func newRouter(routes []RoutesJSON) *httprouter.Router{
	router := newMux()

	for _, n := range routes {
		// Get handler and return if not found
		handler := getHandler(n.Handler)
		if handler == nil && n.Method != "SERVE_DIR" {
			log.Fatal("Handler not found: getHandler() failed on path", n.Path)
		}

		switch n.Method {
			case "GET":
			router.GET(n.Path, handler)
			case "POST":
			router.POST(n.Path, handler)
			case "SERVE_DIR":
			router.ServeFiles(n.Path, http.Dir(n.Handler))
		}
	}

	return router
}

func	newMux() *httprouter.Router{
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(handlers.NotFoundHandler)
	mux.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	mux.HandleMethodNotAllowed = true

	return mux
}

func getHandler(handlerName string) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	switch handlerName {
		case "HealthzHandler":
		return handlers.HealthzHandler
		case "PingHandler":
		return handlers.PingHandler
		case "ReverseHandler":
		return handlers.ReverseHandler
		case "IndexHandler":
		return handlers.IndexHandler
		case "GenericHandler":
		return handlers.GenericHandler
		case "ElementsHandler":
		return handlers.ElementsHandler
	}
	return nil
}
