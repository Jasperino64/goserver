package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Jasperino64/goserver/internal/auth"
	"github.com/Jasperino64/goserver/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID  `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email string `json:"email"`
}
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
	fmt.Printf("Creating user with email: %s\n", req.Email)
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