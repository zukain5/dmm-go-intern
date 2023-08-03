package statuses

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handle request for `GET /v1/statuses/{id}`
func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	elems := strings.Split(r.URL.Path, "/")
	id := elems[len(elems)-1]
	status, err := h.sr.Find(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
