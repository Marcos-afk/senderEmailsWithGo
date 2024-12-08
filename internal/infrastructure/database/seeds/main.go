package main

import (
	"senderEmails/internal/domain/user"
	"senderEmails/internal/infrastructure/database"
	"senderEmails/internal/infrastructure/libs"
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)


var users = []user.User{
	{
		ID: xid.New().String(),
		Name: "ADMIN",
		Email: "admin@email.com",
		Password: func() string {
			hashedPassword, err := libs.HashPassword("12345678")
			if err != nil {
				panic(err)
			}
			return hashedPassword
		}(),
		CreatedAt: time.Now(),
	},
}

func main() {

 db := database.NewConnectionToDB()

 db.Transaction(func(tx *gorm.DB) error {
	for _, u := range users {
		if err := tx.Create(&u).Error; err != nil {
			return err
		}
	}
	return nil
 })
}