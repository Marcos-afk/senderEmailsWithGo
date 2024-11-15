package campaign

import (
	internalerrors "senderEmails/internal/internal-errors"
	"strings"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

type Campaign struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"min=5,max=100"`
	Content   string    `json:"content" validate:"min=5,max=200"`
	Contacts  []Contact `json:"contacts" validate:"min=1,dive"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
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

	campaign := &Campaign{
		ID:        guid.String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}

	validateStructError := internalerrors.ValidateStruct(campaign)
	if validateStructError != nil {
		return nil, validateStructError
	}

	return campaign, nil
}
