package mocks

import (
	"senderEmails/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(newCampaign *campaign.Campaign) (*campaign.Campaign, error) {
	args := r.Called(newCampaign)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (r *CampaignRepositoryMock) Get() []campaign.Campaign {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign)
}

func (r *CampaignRepositoryMock) GetById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (r *CampaignRepositoryMock) Update(updateCampaign *campaign.Campaign) (*campaign.Campaign, error) {
	args := r.Called(updateCampaign)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (r *CampaignRepositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}