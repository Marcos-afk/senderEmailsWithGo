package campaign

import (
	"errors"
	"senderEmails/internal/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createCampaign = contracts.CreateCampaign{
		Name: "Test Campaign",
		Content: "Test Content",
		Emails: []string{"teste01@email.com", "teste02@email.com"} ,
	}

	repositoryMock = new (RepositoryMock)
	service = Service{
		Repository: repositoryMock,
	}
)


type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(campaign *Campaign) (Campaign, error){
	args := r.Called(campaign)
	return args.Get(0).(Campaign), args.Error(1)
}

func Test_Create_Campaign(t *testing.T){
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(Campaign{}, nil)
	 
	_, error := service.Create(createCampaign)
	 assert.Nil(error)

}

func Test_Create_Save_Campaign(t *testing.T){
	repositoryMock.On("Create", mock.MatchedBy(
		func(campaign *Campaign) bool {
			if campaign.Name != createCampaign.Name {
				return false
			}

			if campaign.Content != createCampaign.Content {
				return false
			}

			if len(campaign.Contacts) != len(createCampaign.Emails) {
				return false
			}
			
			return true
		},
	)).Return(Campaign{}, nil)

	service.Create(createCampaign)
	
	repositoryMock.AssertExpectations(t)

}


func Test_Create_Campaign_ValidateDomainError(t *testing.T){
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(Campaign{}, nil)

	_, error := service.Create(contracts.CreateCampaign{
		Name: "",
		Content: "",
		Emails: []string{},
	})

	assert.NotNil(error)
	assert.Equal("name is required", error.Error())
}


func Test_Create_Campaign_ValidateRepositoryCreate(t *testing.T){
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(Campaign{}, errors.New("database error"))

	_, error := service.Create(createCampaign)

	assert.Equal("database error", error.Error())

}