package campaign

import (
	"errors"
	"senderEmails/internal/contracts"
	internalerrors "senderEmails/internal/internal-errors"
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

	service = ServiceImp{}
	repositoryMock *RepositoryMock
)


func setUp() {
	repositoryMock = new(RepositoryMock)
	service.Repository = repositoryMock
}

func Test_Create_Campaign(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(&Campaign{}, nil)

	_, err := service.Create(createCampaign)
	assert.Nil(err)
}

func Test_Create_Save_Campaign(t *testing.T) {
	setUp()

	repositoryMock.On("Create", mock.MatchedBy(
		func(campaign *Campaign) bool {
			if campaign.Name != createCampaign.Name {
				return false
			}
			if campaign.Content != createCampaign.Content {
				return false
			}

			if campaign.Status != PendingStatus {
				return false
			}

			if len(campaign.Contacts) != len(createCampaign.Emails) {
				return false
			}
			return true
		},
	)).Return(&Campaign{}, nil)

	service.Create(createCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(&Campaign{}, nil)

	_, err := service.Create(contracts.CreateCampaign{
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

	repositoryMock.On("Create", mock.Anything).Return(&Campaign{}, errors.New("database error"))

	_, error := service.Create(createCampaign)

	assert.Equal(internalerrors.ErrInternal.Error(), error.Error())
}


func Test_GetAll_Campaign(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("Get").Return([]Campaign{})

	foundedCampaigns := service.Get()
	
	assert.NotNil(foundedCampaigns)
}

func Test_GetCampaignById(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, nil)

	_, err := service.GetById("123")
	
	assert.Nil(err)
}


func Test_GetCampaignById_Return_NotFoundError(t *testing.T){
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, errors.New("campanha não encontrada"))

	_, err := service.GetById("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}



func Test_DeleteCampaign_WithSuccess(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, nil)
	repositoryMock.On("Delete", mock.Anything).Return(nil)

	err := service.Delete("123")

	assert.Nil(err)
}

func Test_DeleteCampaign_NotFoundError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, errors.New("campanha não encontrada"))
	repositoryMock.On("Delete", mock.Anything).Return(nil)

	err := service.Delete("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}


func Test_DeleteCampaign_RepositoryError(t *testing.T){
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, nil)
	repositoryMock.On("Delete", mock.Anything).Return(errors.New("database error"))

	err := service.Delete("123")

	assert.Equal("erro ao deletar campanha " + "database error", err.Error())
}

func Test_CancelCampaign_WithSuccess(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{Status: PendingStatus}, nil)
	repositoryMock.On("Update", mock.Anything).Return(&Campaign{}, nil)

	err := service.Cancel("123")

	assert.Nil(err)
}


func Test_CancelCampaign_NotFoundError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{}, errors.New("campanha não encontrada"))

	err := service.Cancel("123")

	assert.NotNil(err)
	assert.Equal("campanha não encontrada", err.Error())
}


func Test_CancelCampaign_CampaignNotPendingError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{Status: SentStatus}, nil)

	err := service.Cancel("123")

	assert.NotNil(err)
	assert.Equal("campanha não pode ser cancelada", err.Error())
}

func Test_CancelCampaign_RepositoryError(t *testing.T) {
	setUp()

	assert := assert.New(t)

	repositoryMock.On("GetById", mock.Anything).Return(&Campaign{Status: PendingStatus}, nil)
	repositoryMock.On("Update", mock.Anything).Return(&Campaign{}, errors.New("database error"))

	err := service.Cancel("123")

	assert.Equal("erro ao cancelar campanha " + "database error", err.Error())
}