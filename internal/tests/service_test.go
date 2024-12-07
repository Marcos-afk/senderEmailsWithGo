package tests

import (
	"errors"
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"
	internalerrors "senderEmails/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)




func Test_Create_Campaign(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("Create", mock.Anything).Return(&campaign.Campaign{}, nil)

	_, err := campaignServiceImp.Create(createCampaign)
	assert.Nil(err)
}

func Test_Create_Save_Campaign(t *testing.T) {
	setUp()

	campaignRepositoryMock.On("Create", mock.MatchedBy(
		func(newCampaign *campaign.Campaign) bool {
			if newCampaign.Name != createCampaign.Name {
				return false
			}
			if newCampaign.Content != createCampaign.Content {
				return false
			}

			if newCampaign.Status != campaign.PendingStatus {
				return false
			}

			if len(newCampaign.Contacts) != len(createCampaign.Emails) {
				return false
			}
			return true
		},
	)).Return(&campaign.Campaign{}, nil)

	campaignServiceImp.Create(createCampaign)

	campaignRepositoryMock.AssertExpectations(t)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("Create", mock.Anything).Return(&campaign.Campaign{}, nil)

	_, err := campaignServiceImp.Create(contracts.CreateCampaign{
		Name:    "",
		Content: "",
		Emails:  []string{},
	})

	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())
}

func Test_Create_Campaign_ValidateRepositoryCreate(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("Create", mock.Anything).Return(&campaign.Campaign{}, errors.New("database error"))

	_, error := campaignServiceImp.Create(createCampaign)

	assert.Equal(internalerrors.ErrInternal.Error(), error.Error())
}


func Test_GetAll_Campaign(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("Get").Return([]campaign.Campaign{})

	foundedCampaigns := campaignServiceImp.Get()
	
	assert.NotNil(foundedCampaigns)
}

func Test_GetCampaignById(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, nil)

	_, err := campaignServiceImp.GetById("123")
	
	assert.Nil(err)
}


func Test_GetCampaignById_Return_NotFoundError(t *testing.T){
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, errors.New("campanha não encontrada"))

	_, err := campaignServiceImp.GetById("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}



func Test_DeleteCampaign_WithSuccess(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, nil)
	campaignRepositoryMock.On("Delete", mock.Anything).Return(nil)

	err := campaignServiceImp.Delete("123")

	assert.Nil(err)
}

func Test_DeleteCampaign_NotFoundError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, errors.New("campanha não encontrada"))
	campaignRepositoryMock.On("Delete", mock.Anything).Return(nil)

	err := campaignServiceImp.Delete("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}


func Test_DeleteCampaign_RepositoryError(t *testing.T){
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, nil)
	campaignRepositoryMock.On("Delete", mock.Anything).Return(errors.New("database error"))

	err := campaignServiceImp.Delete("123")

	assert.Equal("erro ao deletar campanha " + "database error", err.Error())
}

func Test_CancelCampaign_WithSuccess(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{Status: campaign.PendingStatus}, nil)
	campaignRepositoryMock.On("Update", mock.Anything).Return(&campaign.Campaign{}, nil)

	err := campaignServiceImp.Cancel("123")

	assert.Nil(err)
}


func Test_CancelCampaign_NotFoundError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{}, errors.New("campanha não encontrada"))

	err := campaignServiceImp.Cancel("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}


func Test_CancelCampaign_CampaignNotPendingError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{Status: campaign.SentStatus}, nil)

	err := campaignServiceImp.Cancel("123")

	assert.NotNil(err)
	assert.Equal("campanha não pode ser cancelada", err.Error())
}

func Test_CancelCampaign_RepositoryError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignRepositoryMock.On("GetById", mock.Anything).Return(&campaign.Campaign{Status: campaign.PendingStatus}, nil)
	campaignRepositoryMock.On("Update", mock.Anything).Return(&campaign.Campaign{}, errors.New("database error"))

	err := campaignServiceImp.Cancel("123")

	assert.Equal("erro ao cancelar campanha " + "database error", err.Error())
}