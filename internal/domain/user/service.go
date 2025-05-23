package user

import (
	"errors"
	"senderEmails/internal/contracts"
	"senderEmails/internal/infrastructure/providers"
)

type Service interface {
	Login(data contracts.UserLoginRequest) (*contracts.UserLoginResponse, error)
	ValidateUserId(userId string) (bool, error)
}


type ServiceImp struct {
	Repository Repository
	HashProvider providers.HashProvider
	AuthProvider providers.AuthProvider
}


func (s *ServiceImp) Login(data contracts.UserLoginRequest) (*contracts.UserLoginResponse, error) {
	userFound, err := s.Repository.GetByEmail(data.Email)

	if err != nil {
		return nil, errors.New("email e/ou senha inválidos")
	}

	passwordMatch := s.HashProvider.VerifyPassword(data.Password, userFound.Password)
	if !passwordMatch {
		return nil, errors.New("email e/ou senha inválidos")
	}

	token, err := s.AuthProvider.CreateToken(contracts.CreateToken{
		Sub: userFound.ID,
		Name: userFound.Name,
	})

	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	userLogin := contracts.UserLoginResponse{
		Token: token,	
	}

	return &userLogin, nil
}

func (s *ServiceImp) ValidateUserId(userId string) (bool, error) {
	_, err := s.Repository.GetById(userId)
	if err != nil {
		return false, nil
	}
	return true, nil
}