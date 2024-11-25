package endpoints

import (
	"net/http"
	"senderEmails/internal/contracts"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error){
	var request contracts.CreateCampaign

	render.DecodeJSON(r.Body, &request)

	createdCampaign, serviceErr := h.CampaignService.Create(request)

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