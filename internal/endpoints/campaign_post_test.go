package endpoints

import (
	"net/http"
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	body              = contracts.CreateCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
)

func TestCampaignPost(t *testing.T) {
	setUp()

	service.On("Create", mock.MatchedBy(func(request contracts.CreateCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return(&campaign.Campaign{}, nil)
 

	expectedResponse := struct {
		Message  string            `json:"message"`
		ID 		 string            `json:"id"`
	}{
		Message:  "Campanha criada com sucesso!",
		ID: "",
	}

	req, rr := newHttpTest("POST", "/", body)

	response, status, err := handler.CampaignPost(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, status)
	assert.Equal(t, expectedResponse, response)

}