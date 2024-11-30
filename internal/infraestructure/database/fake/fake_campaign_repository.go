package fake

import "senderEmails/internal/domain/campaign"

type FakeCampaignRepository struct{
	campaigns []campaign.Campaign
}


func (c *FakeCampaignRepository) Create(campaign *campaign.Campaign) (campaign.Campaign, error) {
	c.campaigns = append(c.campaigns, *campaign)
	
	return *campaign, nil
}

func (c *FakeCampaignRepository) Get() []campaign.Campaign {
	return c.campaigns
}

func (c *FakeCampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	for _, campaign := range c.campaigns {
		if campaign.ID == id {
			return &campaign, nil
		}
	}

	return &campaign.Campaign{}, nil
}