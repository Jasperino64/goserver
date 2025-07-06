package main

import (
	"encoding/json"
	"net/http"

	"github.com/Jasperino64/goserver/internal/auth"
	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)


func (config *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		User
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	dbuser, err := config.dbQueries.CreateUser(r.Context(), database.CreateUserParams{
		Email:    req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        dbuser.ID,
			CreatedAt: dbuser.CreatedAt,
			UpdatedAt: dbuser.UpdatedAt,
			Email:    dbuser.Email,
		},
	})
}

func (config *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
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
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		User
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	userDb, err := config.dbQueries.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:              userId,
		Email:           req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        userDb.ID,
			Email:    userDb.Email,
		},
	})
}