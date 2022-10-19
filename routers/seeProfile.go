package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jalamar/clone-twitter/bd"
)

func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		JSONError(w, "id is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)
	if err != nil {
		JSONError(w, "error trying find profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
