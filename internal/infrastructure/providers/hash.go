package providers

import "golang.org/x/crypto/bcrypt"

type HashProvider interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

type HashProviderImp struct{}

func (h *HashProviderImp) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func (h *HashProviderImp) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}