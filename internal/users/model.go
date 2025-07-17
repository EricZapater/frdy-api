package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Email		   string `json:"email" db:"email"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	IsActive bool `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}