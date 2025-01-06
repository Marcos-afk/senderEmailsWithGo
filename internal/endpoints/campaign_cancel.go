package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaignId := chi.URLParam(r, "id")

	err := h.CampaignService.Cancel(campaignId)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Campanha cancelada com sucesso!",
	}

	return response, http.StatusOK, err
}