package campaign

import (
	"strings"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCampaign(name string, content string, emails []string) *Campaign {
	guid := xid.New()
	contacts := make([]Contact, 0, len(emails))

	for _, email := range emails {
		contact := Contact{
			ID:        guid.String(),
			Name:      strings.Split(email, "@")[0],
			Email:     email,
			CreatedAt: time.Now(),
		}

		contacts = append(contacts, contact)
	}

	return &Campaign{
		ID:        guid.String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}
}
