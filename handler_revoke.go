package main

import (
	"net/http"

	"github.com/Jasperino64/goserver/internal/auth"
)

func (config *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	_, err = config.dbQueries.RevokeRefreshToken(r.Context(), token)
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}