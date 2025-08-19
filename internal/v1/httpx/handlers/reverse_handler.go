package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/core"
)

type Message struct {
	Input string `json:"input"`
}

func ReverseHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	// Enforce Content-Type: application/json
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		APIErrorJSON(w, "Bad Request", http.StatusBadRequest, "Content-Type must be application/json")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var msg Message
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&msg)
	if err != nil {
		APIErrorJSON(w, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"input": msg.Input, "output": core.ReverseRunes(msg.Input)})
}
