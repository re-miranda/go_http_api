package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "/Users/remiranda/go/go_http_api/public/index.html")
}

func GenericHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "/Users/remiranda/go/go_http_api/public/generic.html")
}

func ElementsHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "/Users/remiranda/go/go_http_api/public/elements.html")
}
