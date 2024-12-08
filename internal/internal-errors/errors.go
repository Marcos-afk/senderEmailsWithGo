package internalerrors

import (
	"errors"
	"net/http"
	"strings"
)

var ErrInternal error = errors.New("internal server error")


func GetStatusCodeFromError(err error) int {
	errorHandlers := map[string]func(error) int{
		"internal": func(err error) int { return http.StatusInternalServerError },
		"not_found": func(err error) int { return http.StatusNotFound },
		"invalid_credentials": func(err error) int { return http.StatusUnauthorized },
	}

	if errors.Is(err, ErrInternal) {
		return errorHandlers["internal"](err)
	}

	if isNotFoundError(err) {
		return errorHandlers["not_found"](err)
	}

	if isInvalidCredentialsError(err) {
		return errorHandlers["invalid_credentials"](err)
	}

	return http.StatusBadRequest
}



func isNotFoundError(err error) bool {
	notFoundErrors := []string{
		"não encontrada",
		"não encontrado",
	}
	
	for _, message := range notFoundErrors {
		if strings.Contains(err.Error(), message) {
			return true
		}
	}

	return false
}

func isInvalidCredentialsError(err error) bool {
 invalidCredentialsErrors := []string{
		"email e/ou senha inválidos",
 }
	
 for _, message := range invalidCredentialsErrors {
	 if strings.Contains(err.Error(), message) {
		 return true
	 }
 }

 return false
}