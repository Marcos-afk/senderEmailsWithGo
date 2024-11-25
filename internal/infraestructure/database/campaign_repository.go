package database

import "senderEmails/internal/domain/campaign"

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