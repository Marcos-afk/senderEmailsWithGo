package endpoints

import "senderEmails/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}