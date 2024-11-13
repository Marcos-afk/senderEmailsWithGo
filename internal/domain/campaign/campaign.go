package campaign

import (
	"errors"
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
	Name      string    `json:"name" `
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	guid := xid.New()
	contacts := make([]Contact, 0, len(emails))

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("contacts are required")
	}

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
	}, nil
}
