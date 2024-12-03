package endpoints

import (
	"errors"
	"net/http"
	internalerrors "senderEmails/internal/internal-errors"
	"strings"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		var statusCode int

		if err != nil {
			
			if errors.Is(err, internalerrors.ErrInternal){
				statusCode = http.StatusInternalServerError
			} else if strings.Contains(err.Error(), "não encontrada") || strings.Contains(err.Error(), "não encontrado"){
				statusCode = http.StatusNotFound
			}else{
				statusCode = http.StatusBadRequest
			}

			render.Status(r, statusCode)
			render.JSON(w, r, struct {
				Message string `json:"message"`
				Status  int32  `json:"status"`
			}{
				Message: err.Error(),
				Status:  int32(statusCode),
			})
			return
		}

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
			return
		}

		render.NoContent(w, r)
	})
}
