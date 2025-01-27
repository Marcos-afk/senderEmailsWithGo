package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Campaign_Start(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignServiceMock.On("Start", mock.Anything).Return(nil)
	req, rr := newHttpTest("PATCH", "/start", nil)
	req = addParameter(req, "id", "123")

	response, status, err := handler.CampaignStart(rr, req)

	expectedResponse := struct {
		Message string `json:"message"`
	}{
		Message: "Campanha iniciada com sucesso!",
	}

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.Equal(expectedResponse, response)
}