package user

import (
	"senderEmails/internal/domain/campaign"
	"time"
)

type User struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password" validate:"required"`
	Campaigns []campaign.Campaign `json:"campaigns"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}



