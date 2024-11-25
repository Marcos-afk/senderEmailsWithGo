package endpoints

import (
	"errors"
	"net/http"
	internalerrors "senderEmails/internal/internal-errors"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request)(interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      obj, status, err := endpointFunc(w, r)

			if err != nil {
				if errors.Is(err, internalerrors.ErrInternal) {
					render.Status(r, http.StatusInternalServerError)
				} else {
					render.Status(r, http.StatusBadRequest)
				}

				errorResponse := struct {
					Message string `json:"message"`
					Status int32 `json:"status"`
				}{
					Message: err.Error(),
					Status: 400,
				}
		
				render.JSON(w, r, errorResponse)
				return
			}

			render.Status(r, status)
			if obj != nil {
				render.JSON(w, r, obj)
				return
			}
      
			render.Status(r, http.StatusNoContent)
			render.NoContent(w, r)
	})
}