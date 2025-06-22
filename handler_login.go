package main

import (
	"net/http"

	"github.com/Jasperino64/goserver/internal/auth"
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
	
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:    user.Email,
	})

}