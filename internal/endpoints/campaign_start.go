package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaignId := chi.URLParam(r, "id")
  
	err := h.CampaignService.Start(campaignId)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Campanha iniciada com sucesso!",
	}

	return response, http.StatusOK, err
}