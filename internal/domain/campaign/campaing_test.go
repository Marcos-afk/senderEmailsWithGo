package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


var (
	name = "Test Campaign"
	content = "<h1>Test Content</h1>"
	contacts = []string{"contact01@email.com", "contact02@email.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T){
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Len(campaign.Contacts, 2)
} 

func Test_NewCampaign_IdIsNotNil(t *testing.T){
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}


func Test_NewCampaign_CreateAtMustBeNow(t *testing.T){
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	dateNow := time.Now().Add(-time.Minute)

	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, dateNow)
}
