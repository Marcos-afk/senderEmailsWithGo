package endpoints

import (
	"net/http"
	"net/http/httptest"
	"senderEmails/internal/infraestructure/database/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CampaignGetById_ReturnsCampaign(t *testing.T) {
	assert := assert.New(t)

	CampaignServiceConst.Repository = &fake.FakeCampaignRepository{}

	HandlerConst.CampaignService = CampaignServiceConst

	req, err := http.NewRequest("GET", "/campaigns/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	response, status, err := HandlerConst.CampaignGetById(rr, req)

	assert.Nil(err)
	assert.Equal(http.StatusOK, status)
	assert.NotNil(response)
}