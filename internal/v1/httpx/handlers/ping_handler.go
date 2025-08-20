package handlers

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func PingHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "pong")
}
