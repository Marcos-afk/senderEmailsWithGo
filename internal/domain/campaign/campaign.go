package campaign

import (
	internalerrors "senderEmails/internal/internal-errors"
	"strings"
	"time"

	"github.com/rs/xid"
)


const (
	PendingStatus = "PENDING"
	SentStatus    = "SENT"
	FailedStatus  = "FAILED"
	CanceledStatus = "CANCELED"
)


type Contact struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	CampaignId string   `json:"campaign_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

type Campaign struct {
	ID        string     `json:"id" validate:"required"`
	Name      string     `json:"name" validate:"min=5,max=100"`
	Content   string     `json:"content" validate:"min=5"`
	Contacts  []Contact  `json:"contacts" validate:"min=1,dive"`
	Status 		string     `json:"status" validate:"required"`
	UserId    string     `json:"user_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}



func NewCampaign(name string, content string, createdBy string, emails []string) (*Campaign, error) {
	campaignId := xid.New().String()
	contacts := make([]Contact, 0, len(emails))

	for _, email := range emails {
		contact := Contact{
			ID:        xid.New().String(),
			Name:      strings.Split(email, "@")[0],
			CampaignId: campaignId,
			Email:     email,
			CreatedAt: time.Now(),
		}

		contacts = append(contacts, contact)
	}

	campaign := &Campaign{
		ID:        campaignId,
		Name:      name,
		Content:   content,
		Status: 	PendingStatus,
		UserId: createdBy,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}

	validateStructError := internalerrors.ValidateStruct(campaign)
	if validateStructError != nil {
		return nil, validateStructError
	}

	return campaign, nil
}


func (c *Campaign) Cancel() {
	c.Status = CanceledStatus
}


func (c *Campaign) Sent(){
	c.Status = SentStatus
}