package campaign

import (
	"time"

	"github.com/google/uuid"
)


type Contact struct {
	ID        string `json:"id"`
	Name			string `json:"name"`
	Email			string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Campaign struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Contacts  []Contact `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}



func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))

	for _, email := range emails {
		contact := Contact{
			ID:        uuid.New().String(),
			Name:      "Name",
			Email:     email,
			CreatedAt: time.Now(),
		}
		contacts = append(contacts, contact)
	}

	return &Campaign{
		ID:        uuid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}
}