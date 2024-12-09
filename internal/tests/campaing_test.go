package tests

import (
	"senderEmails/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func Test_NewCampaign_CreateCampaign(t *testing.T){
	assert := assert.New(t)

	createdCampaign, _ := campaign.NewCampaign(name, content,createdBy, contacts)

	assert.Equal(name, createdCampaign.Name)
	assert.Equal(campaign.PendingStatus, createdCampaign.Status)
	assert.Equal(content, createdCampaign.Content)
	assert.Equal(createdBy, createdCampaign.UserId)
	assert.Len(createdCampaign.Contacts, 2)
} 

func Test_NewCampaign_IdIsNotNil(t *testing.T){
	assert := assert.New(t)

	createdCampaign, _ := campaign.NewCampaign(name, content,createdBy, contacts)

	assert.NotNil(createdCampaign.ID)
}


func Test_NewCampaign_CreateAtMustBeNow(t *testing.T){
	assert := assert.New(t)

	createdCampaign, _ := campaign.NewCampaign(name, content,createdBy, contacts)

	dateNow := time.Now().Add(-time.Minute)

	assert.NotNil(createdCampaign.CreatedAt)
	assert.Greater(createdCampaign.CreatedAt, dateNow)
}


func Test_NewCampaign_NameNotEmpty(t *testing.T){
	assert := assert.New(t)

	_, err := campaign.NewCampaign("", content,createdBy, contacts)

	assert.Equal("name is required with min 5", err.Error())
}


func Test_NewCampaign_ValidateNameMax(t *testing.T){
	assert := assert.New(t)

	_, err := campaign.NewCampaign(fake.Lorem().Text(300), content,createdBy, contacts)

	assert.Equal("name is required with max 100", err.Error())
}


func Test_NewCampaign_ContactsNotEmpty(t *testing.T){
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, createdBy, []string{})

	assert.Equal("contacts is required with min 1", err.Error())
}
