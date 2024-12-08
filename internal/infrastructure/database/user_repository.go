package database

import (
	"errors"
	"senderEmails/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (u *UserRepository) GetByEmail(email string) (*user.User, error) {
	var foundUser user.User

	tx := u.Db.Where("email = ?", email).First(&foundUser)

	if tx.Error != nil {
		return &user.User{}, errors.New("usuário não encontrado")
	}

	return &foundUser, nil
}


func (u *UserRepository) GetById(id string) (*user.User, error) {
	var foundUser user.User

	tx := u.Db.Where("id = ?", id).First(&foundUser)

	if tx.Error != nil {
		return &user.User{}, errors.New("usuário não encontrado")
	}

	return &foundUser, nil
}