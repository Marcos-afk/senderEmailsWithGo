package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"senderEmails/internal/infrastructure/providers"
	"senderEmails/internal/tests/mocks"

	"github.com/go-chi/chi/v5"
)


func setUp(){
	campaignRepositoryMock = new(mocks.CampaignRepositoryMock)
	campaignServiceMock = new (mocks.CampaignServiceMock)
	mailProviderMock  = new (providers.FakeMailProvider)
	campaignServiceImp.Repository = campaignRepositoryMock
	campaignServiceImp.MailProvider = mailProviderMock
	handler.CampaignService = campaignServiceMock
}

func newHttpTest(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {

	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}

	req, _ := http.NewRequest(method, url, &buf)
	rr := httptest.NewRecorder()
	return req, rr
}

func addParameter(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(keyParameter, valueParameter)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}

