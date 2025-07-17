package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {
	type upgradeRequest struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}
	var req upgradeRequest
	if err := parseJSON(r, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	userId, err := uuid.Parse(req.Data.UserID)
	if err != nil || userId == uuid.Nil {
		http.Error(w, "Invalid user ID", http.StatusNotFound)
		return
	}

	_, err = config.dbQueries.SetIsChirpyRed(r.Context(), database.SetIsChirpyRedParams {
		ID:        userId,
		IsChirpyRed: true,
	})
	if err != nil {
		if (errors.Is(err, sql.ErrNoRows)) {
			respondWithError(w, http.StatusNotFound, "User not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}