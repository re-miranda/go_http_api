package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func HealthzHandler (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
