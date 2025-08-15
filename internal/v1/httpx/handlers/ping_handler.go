package handlers

import "net/http"
import "fmt"

func PingHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
