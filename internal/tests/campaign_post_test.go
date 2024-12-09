package tests

import (
	"context"
	"net/http"
	"senderEmails/internal/contracts"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/infrastructure/middlewares"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestCampaignPost(t *testing.T) {
	setUp()

	campaignServiceMock.On("Create", mock.MatchedBy(func(request contracts.CreateCampaign) bool {
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


	userID := "test-user-id"
	
	req, rr := newHttpTest("POST", "/", body)
	ctx := context.WithValue(req.Context(), middlewares.UserIdKey, userID)
	req = req.WithContext(ctx)
	

	response, status, err := handler.CampaignPost(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, status)
	assert.Equal(t, expectedResponse, response)

}