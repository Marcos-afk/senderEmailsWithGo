package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignCancelPatch(t *testing.T) {
	setUp()

	assert := assert.New(t)

	campaignServiceMock.On("Cancel", mock.Anything).Return(nil)
	req, rr := newHttpTest("PATCH", "/", nil)
	req = addParameter(req, "id", "123")

	response, status, err := handler.CampaignCancelPatch(rr, req)

	expectedResponse := struct {
		Message string `json:"message"`
	}{
		Message: "Campanha cancelada com sucesso!",
	}

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.Equal(expectedResponse, response)
}