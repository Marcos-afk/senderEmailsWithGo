package database

import (
	"senderEmails/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionToDB() *gorm.DB {
	dbAddress := "host=localhost user=postgres password=postgres dbname=sender_emails port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Falha ao se conectar ao banco de dados! " + err.Error())
	}

	 
	 migrateError := db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	 if migrateError != nil {
		panic("Falha ao criar as tabelas! " + migrateError.Error())
	 }

	
	return db
}