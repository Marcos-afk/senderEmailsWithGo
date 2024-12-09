package endpoints

import (
	"net/http"
	"senderEmails/internal/contracts"
	internalerrors "senderEmails/internal/internal-errors"

	"github.com/go-chi/render"
)

func (h *Handler) UserLoginPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contracts.UserLoginRequest

	render.DecodeJSON(r.Body, &request)

	validateStructError := internalerrors.ValidateStruct(request)
	if validateStructError != nil {
		return nil, http.StatusBadRequest, validateStructError
	}

	session, err := h.UserService.Login(request)

	response := struct {
		Message   string `json:"message"`
		Token      string `json:"token"`
	}{
		Message: "Login realizado com sucesso!",
		Token:      "",
	}

	
	if err == nil {
		response.Token = session.Token
	}

	return response, http.StatusOK, err
}