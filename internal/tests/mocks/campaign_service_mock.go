package mocks

import (
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}


func (c *CampaignServiceMock) Get() []campaign.Campaign {
	args := c.Called()
	return args.Get(0).([]campaign.Campaign)
}

func (c *CampaignServiceMock) GetById(id string) (*contracts.GetCampaignByIdResponse, error) {
	args := c.Called(id)
	return args.Get(0).(*contracts.GetCampaignByIdResponse), args.Error(1)
}

func (c *CampaignServiceMock) Create(newCampaign contracts.CreateCampaign) (*campaign.Campaign, error) {
	args := c.Called(newCampaign)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (c *CampaignServiceMock) Cancel(id string) error {
	args := c.Called(id)
	return args.Error(0)
}

func (c *CampaignServiceMock) Delete(id string) error {
	args := c.Called(id)
	return args.Error(0)
}


func (c *CampaignServiceMock) Start(id string) error {
	args := c.Called(id)
	return args.Error(0)
}


func (c *CampaignServiceMock) SendMailAndUpdateStatus(SendMailToCampaign *campaign.Campaign)(bool, error) {
	args := c.Called(SendMailToCampaign)
	return args.Get(0).(bool), args.Error(1)
}
