package database

import (
	"errors"
	"senderEmails/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct{
	Db *gorm.DB
}


func (c *CampaignRepository) Create(campaignData *campaign.Campaign) (campaign.Campaign, error) {
	tx := c.Db.Create(&campaignData)

	if tx.Error != nil {
		return campaign.Campaign{}, tx.Error
	}
	
	return *campaignData, nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	var campaigns []campaign.Campaign
  
	c.Db.Preload("Contacts").Find(&campaigns)

	return campaigns
}


func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var foundCampaign campaign.Campaign

	tx := c.Db.Where("id = ?", id).First(&foundCampaign)
	if tx.Error != nil {
		return &campaign.Campaign{}, errors.New("campanha não encontrada")
	}

	return &foundCampaign, nil
}