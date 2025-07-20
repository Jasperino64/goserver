package main

import (
	"net/http"
	"sort"

	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)
func (config *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	authorUserId := r.URL.Query().Get("author_id")
	sortVal := r.URL.Query().Get("sort")
	var chirps []database.Chirp
	var err error
	if authorUserId != "" {
		uuidVal, err := uuid.Parse(authorUserId)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID format", err)
			return
		}
		chirps, err = config.dbQueries.GetChirpsByUserID(r.Context(), uuidVal)
		
	} else {
		chirps, err = config.dbQueries.GetAllChirps(r.Context())
	}
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve chirps", err)
		return
	}
	if sortVal == "desc" {
		sortChirps(chirps, true)
	} else if sortVal == "asc" {
		sortChirps(chirps, false)
	}
	
	var response []Chirp
	for _, chirp := range chirps {
		response = append(response, Chirp{
			Id:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserId:    chirp.UserID,
			Body:      chirp.Body,
		})
	}
	respondWithJSON(w, http.StatusOK, response)
}

func sortChirps(chirps []database.Chirp, desc bool) {
	sort.Slice(chirps, func(i, j int) bool {
		if desc {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})
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
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserId:    chirp.UserID,
		Body:      chirp.Body,
	})
}