package user

import (
	"time"
)

type User struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}



