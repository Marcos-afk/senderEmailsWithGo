package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"senderEmails/internal/contracts"
	"senderEmails/internal/infraestructure/database/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_CampaignPost_New_Campaign(t *testing.T){
	assert := assert.New(t)

	CampaignServiceConst.Repository = &fake.FakeCampaignRepository{}

	HandlerConst.CampaignService = CampaignServiceConst

	body := contracts.CreateCampaign{
		Name: "Teste Campaign",
		Content: "Teste Content",
		Emails: []string{"test@email.com"},
	}

	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(body)

	req, err := http.NewRequest("POST", "/", &buffer)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	response, status, err := HandlerConst.CampaignPost(rr, req)
	assert.Nil(err)
	assert.Equal(http.StatusCreated, status)
	assert.NotNil(response)
}