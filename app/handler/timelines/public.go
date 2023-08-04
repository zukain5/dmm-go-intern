package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p := object.NewFindPublicTimelinesParams(r.URL.Query())
	timeline, err := h.tr.FindPublicTimelines(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
