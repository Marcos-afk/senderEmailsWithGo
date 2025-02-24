package main

import (
	"fmt"
	"senderEmails/internal/infrastructure/database"
)

func main() {
	db := database.NewConnectionToDB()

	campaignsRepository := database.CampaignRepository{
		Db: db,
	}

	campaigns := campaignsRepository.Get()

	fmt.Println(campaigns)
}