package tests

import (
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/endpoints"
	"senderEmails/internal/tests/mocks"

	"github.com/jaswdr/faker/v2"
)

var (
	name     = "Test Campaign"
	content  = "<h1>Test Content</h1>"
	contacts = []string{"contact01@email.com", "contact02@email.com"}
	fake     = faker.New()
	body     = contracts.CreateCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	createCampaign = contracts.CreateCampaign{
		Name:    name,
		Content: content,
		Emails:  contacts,

	}
	campaignRepositoryMock *mocks.CampaignRepositoryMock
	campaignServiceMock *mocks.CampaignServiceMock
	campaignServiceImp  = &campaign.ServiceImp{}
	handler = endpoints.Handler{}
)