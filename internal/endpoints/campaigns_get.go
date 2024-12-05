package endpoints

import (
	"net/http"
	"senderEmails/internal/domain/campaign"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaignsFound := h.CampaignService.Get()

	response := struct {
		Message  string              `json:"message"`
		Campaigns []campaign.Campaign `json:"campaigns"`
	}{
		Message:  "Campanhas encontradas com sucesso!",
		Campaigns: campaignsFound,
	}

	return response, http.StatusOK, nil
}