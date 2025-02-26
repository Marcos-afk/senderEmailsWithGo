package main

import (
	"log"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/infrastructure/database"
	"senderEmails/internal/infrastructure/providers"
	"time"
)

func main() {
	db := database.NewConnectionToDB()

	campaignsRepository := database.CampaignRepository{
		Db: db,
	}

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{
			Db: db,
		},
		MailProvider: &providers.MailProviderImp{},
	}

	for {
		campaigns := campaignsRepository.GetCampaignsToBeSent()
	
		for _, campaign := range campaigns {
			success, error := campaignService.SendMailAndUpdateStatus(&campaign); if error != nil {
				log.Println(error)
			}

			if success {
				log.Println("Email's da campanha '" + campaign.Name + "' enviados com sucesso")
			}
		}

		time.Sleep(20 * time.Minute)
	}

}