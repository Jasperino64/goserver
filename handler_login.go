package main

import (
	"net/http"
	"time"

	"github.com/Jasperino64/goserver/internal/auth"
)

func (config *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		ExpiresInSeconds int64 `json:"expires_in_seconds"`
	}
	err := parseJSON(r, &req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if req.Email == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Email and password are required", nil)
		return
	}
	user, err := config.dbQueries.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email", nil)
		return
	}
	err = auth.CheckPasswordHash(req.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}
	
	if req.ExpiresInSeconds <= 0 {
		req.ExpiresInSeconds = 3600 // Default to 1 hour if not specified
	}
	token, err := auth.MakeJWT(user.ID, config.secretKey, time.Duration(req.ExpiresInSeconds)*time.Second)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create JWT token", err)
		return
	}
	w.Header().Set("Authorization", "Bearer " + token)
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:    user.Email,
		Token:  token,
	})

}