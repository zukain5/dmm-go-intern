package accounts

import (
	"encoding/json"
	"net/http"
	"strings"
	"yatter-backend-go/app/handler/auth"
)

// Handle request for `POST /v1/accounts/{username}/follow`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	elems := strings.Split(r.URL.Path, "/")
	username := elems[len(elems)-2]

	follower := auth.AccountOf(r)
	followee, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if followee == nil {
		http.Error(w, "Follow target account not found.", http.StatusNotFound)
		return
	}

	id, followed_by, err := h.rr.Create(ctx, follower, followee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res := struct {
		ID         int64 `json:"id"`
		Following  bool  `json:"following"`
		FollowedBy bool  `json:"followed_by"`
	}{
		ID:         id,
		Following:  true,
		FollowedBy: followed_by,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
