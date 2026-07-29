package handler

import (
	"encoding/json"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"version": "0.0.1",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
