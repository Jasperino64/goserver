package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		http.Error(w, "Not allowed", http.StatusForbidden)
		return
	}
	cfg.fileServerHits.Store(0)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileServerHits.Load())))


	err := cfg.dbQueries.DeleteAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}