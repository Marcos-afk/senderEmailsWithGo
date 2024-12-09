package endpoints

import (
	"net/http"
	"senderEmails/internal/contracts"
	"senderEmails/internal/infrastructure/middlewares"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error){
	var request contracts.CreateCampaignRequest

	render.DecodeJSON(r.Body, &request)

	createdBy := r.Context().Value(middlewares.UserIdKey).(string)

	createdCampaign, serviceErr := h.CampaignService.Create(contracts.CreateCampaign{
		Name: request.Name,
		Content: request.Content,
		CreatedBy: createdBy,
		Emails: request.Emails,
	})

	response := struct {
		Message string `json:"message"`
		ID      string `json:"id"`
	}{
		Message: "Campanha criada com sucesso!",
		ID:      "",
	}

	if serviceErr == nil {
		response.ID = createdCampaign.ID
	}
	
	return response, http.StatusCreated, serviceErr
}