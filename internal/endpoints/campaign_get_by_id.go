package endpoints

import (
	"net/http"
	"senderEmails/internal/contracts"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaignId := chi.URLParam(r, "id")
	campaignFound, err := h.CampaignService.GetById(campaignId)

	response := struct {
		Message  string            `json:"message"`
		Campaign contracts.GetCampaignByIdResponse `json:"campaign"`
	}{
		Message:  "Campanha encontrada com sucesso!",
		Campaign: campaignFound,
	}

	return response, http.StatusOK, err
}