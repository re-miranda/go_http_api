package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/re-miranda/go_http_api/internal/v1/core"
	"github.com/julienschmidt/httprouter"
)

func ReverseHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		APIErrorJSON(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Message struct {
		Input string `json:"input"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var m Message
	err := dec.Decode(&m)
	if err != nil {
		APIErrorJSON(w, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"input": m.Input, "output": core.ReverseRunes(m.Input)})
}
