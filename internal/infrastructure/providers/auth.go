package providers

import (
	"errors"
	"senderEmails/internal"
	"senderEmails/internal/contracts"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(internal.JWT_SECRET_KEY)

type AuthProvider interface {
	CreateToken(createToken contracts.CreateToken) (string, error)
	VerifyToken(tokenString string, validateUserFunc func(userId string) (bool, error)) (string, error)
}

type AuthProviderImp struct {}

func (a *AuthProviderImp) CreateToken(createToken contracts.CreateToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": 		createToken.Sub,
			"name": 	createToken.Name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthProviderImp) VerifyToken(tokenString string, validateUserFunc func(userId string) (bool, error)) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	sub, valid := claims["sub"].(string)
	if !ok || !valid || sub == "" {
		return "", errors.New("token inválido")
	}

	isValid, err := validateUserFunc(sub)
	if err != nil || !isValid {
		return "", errors.New("token inválido")
	}

	return sub, nil
}