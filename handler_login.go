package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Jasperino64/goserver/internal/auth"
	"github.com/Jasperino64/goserver/internal/database"
)

func (config *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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
	
	token, err := auth.MakeJWT(
		user.ID,
		config.secretKey,
		time.Hour,
	)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create JWT token", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()

	config.dbQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID: user.ID,
		Token: refreshToken,
		ExpiresAt: time.Now().Add(60 * 24 * time.Hour).UTC(), // refresh token valid for 60 days
		RevokedAt: sql.NullTime{},
	})
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create refresh token", err)
		return
	}
	w.Header().Set("Authorization", "Bearer " + token)
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:    user.Email,
		Token:  token,
		RefreshToken: refreshToken,
		IsChirpyRed: user.IsChirpyRed,
	})

}