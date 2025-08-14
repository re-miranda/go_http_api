// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package reverse

import (
	"encoding/json"
	"net/http"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func reverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func ReverseHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	if !dec.More() {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	type Message struct {
		Input string `json:"input"`
	}

	dec.DisallowUnknownFields()
	var m Message
	err := dec.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"input": m.Input, "output": reverseRunes(m.Input)})
}
