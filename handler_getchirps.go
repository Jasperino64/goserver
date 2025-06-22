package main

import (
	"net/http"

	"github.com/google/uuid"
)
func (config *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := config.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve chirps", err)
		return
	}
	var response []Chirp
	for _, chirp := range chirps {
		response = append(response, Chirp{
			Id:        chirp.ID,
			CreatedAt: chirp.CreatedAt.String(),
			UpdatedAt: chirp.UpdatedAt.String(),
			UserId:    chirp.UserID,
			Body:      chirp.Body,
		})
	}
	respondWithJSON(w, http.StatusOK, response)
}

func (config *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirp_id")
	if chirpId == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", nil)
		return
	}
	uuidVal, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID format", err)
		return
	}
	chirp, err := config.dbQueries.GetChirpByID(r.Context(), uuidVal)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, Chirp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt.String(),
		UpdatedAt: chirp.UpdatedAt.String(),
		UserId:    chirp.UserID,
		Body:      chirp.Body,
	})
}