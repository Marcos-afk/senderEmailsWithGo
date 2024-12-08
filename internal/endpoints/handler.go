package endpoints

import (
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/domain/user"
)

type Handler struct {
	CampaignService campaign.Service
	UserService     user.Service
}