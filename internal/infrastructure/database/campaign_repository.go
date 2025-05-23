package database

import (
	"errors"
	"senderEmails/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct{
	Db *gorm.DB
}


func (c *CampaignRepository) Create(campaignData *campaign.Campaign) (*campaign.Campaign, error) {
	tx := c.Db.Create(&campaignData)

	if tx.Error != nil {
		return nil, tx.Error
	}
	
	return campaignData, nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	var campaigns []campaign.Campaign
	
	c.Db.Preload("Contacts").Find(&campaigns)

	return campaigns
}


func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var foundCampaign campaign.Campaign

	tx := c.Db.Preload("Contacts").Where("id = ?", id).First(&foundCampaign)
	if tx.Error != nil {
		return &campaign.Campaign{}, errors.New("campanha não encontrada")
	}

	return &foundCampaign, nil
}


func (c *CampaignRepository) Update(campaignData *campaign.Campaign) (*campaign.Campaign, error) {
	tx := c.Db.Save(&campaignData)

	if tx.Error != nil {
		return nil, tx.Error
	}
	
	return campaignData, nil
}

func (c *CampaignRepository) Delete(id string) error {
	tx := c.Db.Begin()

	if err := tx.Where("campaign_id = ?", id).Delete(&campaign.Contact{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&campaign.Campaign{}).Error; err != nil {
		tx.Rollback() 
		return err
	}

	return tx.Commit().Error
}


func (c *CampaignRepository) GetCampaignsToBeSent() []campaign.Campaign {
	var campaigns []campaign.Campaign

	c.Db.Preload("Contacts").
		Where("status = ? and date_part('minute', now()::timestamp - updated_at::timestamp) >= ? ",
		campaign.StartedStatus, 5).
	 	Find(&campaigns)

	return campaigns
}