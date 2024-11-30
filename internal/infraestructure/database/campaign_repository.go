package database

import (
	"errors"
	"senderEmails/internal/domain/campaign"
)

type CampaignRepository struct{
	campaigns []campaign.Campaign
}


func (c *CampaignRepository) Create(campaign *campaign.Campaign) (campaign.Campaign, error) {
	c.campaigns = append(c.campaigns, *campaign)
	
	return *campaign, nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	return c.campaigns
}


func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	for _, campaign := range c.campaigns {
		if campaign.ID == id {
			return &campaign, nil
		}
	}

	return &campaign.Campaign{}, errors.New("campanha não encontrada")
}