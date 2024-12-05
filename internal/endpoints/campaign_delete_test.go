package endpoints

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignDelete(t *testing.T) {
	setUp()	


	assert := assert.New(t)

	service.On("Delete", mock.Anything).Return(nil)
	req, rr := newHttpTest("DELETE", "/", nil)
	req = addParameter(req, "id", "123")

	response, status, err := handler.CampaignDelete(rr, req)

	expectedResponse := struct {
		Message string `json:"message"`
	}{
		Message: "Campanha apagada com sucesso!",
	}

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.Equal(expectedResponse, response)
}