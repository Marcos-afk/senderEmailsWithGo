package campaign

import (
	"senderEmails/internal/contracts"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}


func (c *CampaignServiceMock) Get() []Campaign {
	args := c.Called()
	return args.Get(0).([]Campaign)
}

func (c *CampaignServiceMock) GetById(id string) (*contracts.GetCampaignByIdResponse, error) {
	args := c.Called(id)
	return args.Get(0).(*contracts.GetCampaignByIdResponse), args.Error(1)
}

func (c *CampaignServiceMock) Create(newCampaign contracts.CreateCampaign) (*Campaign, error) {
	args := c.Called(newCampaign)
	return args.Get(0).(*Campaign), args.Error(1)
}

func (c *CampaignServiceMock) Cancel(id string) error {
	args := c.Called(id)
	return args.Error(0)
}

func (c *CampaignServiceMock) Delete(id string) error {
	args := c.Called(id)
	return args.Error(0)
}
