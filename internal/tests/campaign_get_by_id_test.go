package tests

import (
	"net/http"
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignId := "123"
	campaignResponse := contracts.GetCampaignByIdResponse {
		ID: campaignId,
		Name: "Campaign 1",
		Content: "Content 1",
		Status: campaign.PendingStatus,
	} 

	campaignServiceMock.On("GetById", mock.Anything).Return(&campaignResponse, nil)
	req, rr := newHttpTest("GET", "/", nil)
	req = addParameter(req, "id", campaignId)

	response, status, err := handler.CampaignGetById(rr, req)

	expectedResponse := struct {
		Message  string                            `json:"message"`
		Campaign contracts.GetCampaignByIdResponse `json:"campaign"`
	}{
		Message:  "Campanha encontrada com sucesso!",
		Campaign: campaignResponse,
	}

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.Equal(expectedResponse, response)

}
