package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)


var (
	name = "Test Campaign"
	content = "<h1>Test Content</h1>"
	contacts = []string{"contact01@email.com", "contact02@email.com"}
	fake = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Len(campaign.Contacts, 2)
} 

func Test_NewCampaign_IdIsNotNil(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}


func Test_NewCampaign_CreateAtMustBeNow(t *testing.T){
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	dateNow := time.Now().Add(-time.Minute)

	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, dateNow)
}


func Test_NewCampaign_NameNotEmpty(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}


func Test_NewCampaign_ValidateNameMax(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(300), content, contacts)

	assert.Equal("name is required with max 100", err.Error())
}

func Test_NewCampaign_ContentNotEmpty(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(250), contacts)

	assert.Equal("content is required with max 200", err.Error())
}

func Test_NewCampaign_ContactsNotEmpty(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required with min 1", err.Error())
}
