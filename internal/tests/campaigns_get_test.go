package tests

import (
	"net/http"
	"senderEmails/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CampaignsGet(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignServiceMock.On("Get").Return([]campaign.Campaign{}, nil)
	req, rr := newHttpTest("GET", "/", nil)

	response, status, err := handler.CampaignsGet(rr, req)

	expectedResponse := struct {
		Message  string              `json:"message"`
		Campaigns []campaign.Campaign `json:"campaigns"`
	}{
		Message:  "Campanhas encontradas com sucesso!",
		Campaigns: []campaign.Campaign{},
	}

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.Equal(expectedResponse, response)

}

