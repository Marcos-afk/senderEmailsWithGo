package middlewares

import (
	"context"
	"net/http"
	"senderEmails/internal/domain/user"
	"senderEmails/internal/infrastructure/providers"

	"github.com/go-chi/render"
)

type AuthMiddlewareResponse struct {
	Message string `json:"message"`
}

type contextKey string

const UserIdKey = contextKey("userId")

var authProvider = providers.AuthProviderImp{}

func AuthMiddleware(userService user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, AuthMiddlewareResponse{Message: "token não informado"})
				return
			}

			prefix := "Bearer "
			if len(tokenString) <= len(prefix) || tokenString[:len(prefix)] != prefix {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, AuthMiddlewareResponse{Message: "token inválido"})
				return
			}

			tokenString = tokenString[len(prefix):]
			sub, err := authProvider.VerifyToken(tokenString, userService.ValidateUserId)
			if err != nil {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, AuthMiddlewareResponse{Message: "token inválido"})
				return
			}


			ctx := context.WithValue(r.Context(), UserIdKey, sub)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}