package main

import (
	"net/http"

	"github.com/Jasperino64/goserver/internal/auth"
	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
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
	var req struct {
		Body    string `json:"body"`
	}
	if err := parseJSON(r, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.Body == "" {
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
		return
	}
	chirp, err := config.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		UserID: userId,
		Body:   req.Body,
	})
	if err != nil {
		http.Error(w, "Failed to create chirp", http.StatusInternalServerError)
		return
	}
	
	respondWithJSON(w, http.StatusCreated, Chirp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserId:    chirp.UserID,
		Body:      chirp.Body,
	})
}
