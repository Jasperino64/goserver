package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jasperino64/goserver/internal/auth"
)

func (config *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	user, err := config.dbQueries.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}
	// Create a new JWT token
	token, err := auth.MakeJWT(user.ID, config.secretKey, 1*time.Hour)
	if err != nil {
		fmt.Println("Error creating JWT token:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create JWT token", err)
		return
	}
	
	respondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}