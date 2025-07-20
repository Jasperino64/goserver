package main

import (
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	Id     uuid.UUID `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UserId uuid.UUID `json:"user_id"`
	Body   string    `json:"body"`
}