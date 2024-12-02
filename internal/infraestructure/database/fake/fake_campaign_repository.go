package fake

import (
	"errors"
	"senderEmails/internal/domain/campaign"
)

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

func (c *FakeCampaignRepository) Update(updatedCampaign *campaign.Campaign) (campaign.Campaign, error) {
	for i, existingCampaign := range c.campaigns {
		if existingCampaign.ID == updatedCampaign.ID {
			c.campaigns[i] = *updatedCampaign
			return c.campaigns[i], nil
		}
	}
	return campaign.Campaign{}, errors.New("campanha não encontrada")
}

func (c *FakeCampaignRepository) Delete(id string) error {
	for i, existingCampaign := range c.campaigns {
		if existingCampaign.ID == id {
			c.campaigns = append(c.campaigns[:i], c.campaigns[i+1:]...)
			return nil
		}
	}
	
	return errors.New("campanha não encontrada")
}