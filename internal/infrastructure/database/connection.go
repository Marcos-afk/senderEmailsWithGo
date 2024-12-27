package database

import (
	"senderEmails/internal"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/domain/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionToDB() *gorm.DB {
	dbAddress := internal.DATABASE_URL

	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Falha ao se conectar ao banco de dados! " + err.Error())
	}

	 
	 migrateError := db.AutoMigrate(&user.User{},&campaign.Campaign{}, &campaign.Contact{})

	 if migrateError != nil {
		panic("Falha ao criar as tabelas! " + migrateError.Error())
	 }

	
	return db
}

