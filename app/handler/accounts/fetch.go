package accounts

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	elems := strings.Split(r.URL.Path, "/")
	username := elems[len(elems)-1]
	account, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
