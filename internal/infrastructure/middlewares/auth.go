package middlewares

import (
	"net/http"
	"senderEmails/internal/domain/user"
	"senderEmails/internal/infrastructure/libs"

	"github.com/go-chi/render"
)

type AuthMiddlewareResponse struct {
	Message string `json:"message"`
}

func AuthMiddleware(userService user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, AuthMiddlewareResponse{Message: "token não informado"})
				return
			}

			err := libs.VerifyToken(tokenString, userService.ValidateUserId)
			if err != nil {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, AuthMiddlewareResponse{Message: "token inválido"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}