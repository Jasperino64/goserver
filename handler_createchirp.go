package main

import (
	"net/http"

	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	Id     uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	UserId uuid.UUID `json:"user_id"`
	Body   string    `json:"body"`
}

func (config *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId  uuid.UUID `json:"user_id"`
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
		UserID: req.UserId,
		Body:   req.Body,
	})
	if err != nil {
		http.Error(w, "Failed to create chirp", http.StatusInternalServerError)
		return
	}
	
	respondWithJSON(w, http.StatusCreated, Chirp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt.String(),
		UpdatedAt: chirp.UpdatedAt.String(),
		UserId:    chirp.UserID,
		Body:      chirp.Body,
	})
}