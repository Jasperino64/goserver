package main

import (
	"net/http"

	"github.com/Jasperino64/goserver/internal/auth"
	"github.com/google/uuid"
)


func (config *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid or missing token", http.StatusUnauthorized)
		return
	}
	userId, err := auth.ValidateJWT(token, config.secretKey)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}
	if userId == uuid.Nil {
		http.Error(w, "Unauthorized: Invalid user ID", http.StatusUnauthorized)
		return
	}

	chirpIdStr := r.PathValue("chirp_id")
	chirpId, err := uuid.Parse(chirpIdStr)
	if err != nil || chirpId == uuid.Nil {
		http.Error(w, "Invalid chirp ID", http.StatusBadRequest)
		return
	}
	chirp , err := config.dbQueries.GetChirpByID(r.Context(), chirpId)
	if err != nil {
		http.Error(w, "Chirp not found", http.StatusNotFound)
		return
	}
	if chirp.UserID != userId {
		http.Error(w, "Unauthorized: You can only delete your own chirps", http.StatusForbidden)
		return
	}
	err = config.dbQueries.DeleteChirpByID(r.Context(), chirpId)
	if err != nil {
		http.Error(w, "Failed to delete chirp", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}