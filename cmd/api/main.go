package main

import (
	"log"
	"net/http"
	"senderEmails/internal"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/domain/user"
	"senderEmails/internal/endpoints"
	"senderEmails/internal/infrastructure/database"
	"senderEmails/internal/infrastructure/middlewares"
	"senderEmails/internal/infrastructure/providers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func  AllowOriginFunc (r *http.Request, origin string) bool {
	whitelist := internal.WHITELIST
	for _, regex := range whitelist  {
		if regex.MatchString(origin) || origin == "" {
			return true
		}

	}

	return false
}


func main() {
	routes := chi.NewRouter()

	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Use(cors.Handler(cors.Options{
    AllowOriginFunc: AllowOriginFunc,
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300,
}))

	db := database.NewConnectionToDB()

	userService := user.ServiceImp{
		Repository: &database.UserRepository{
			Db: db,
		},
		HashProvider: &providers.HashProviderImp{},
		AuthProvider: &providers.AuthProviderImp{},
	}

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{
			Db: db,
		},
		MailProvider: &providers.MailProviderImp{},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
		UserService:     &userService,
	}

	authMiddleware := middlewares.AuthMiddleware(&userService)

	routes.Route("/api/v1", func(r chi.Router) {
		r.Route("/campaigns", func(r chi.Router) {
			r.Use(authMiddleware)
			r.Post("/", endpoints.HandlerError(handler.CampaignPost))
			r.Get("/", endpoints.HandlerError(handler.CampaignsGet))
			r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
			r.Patch("/{id}/cancel", endpoints.HandlerError(handler.CampaignCancelPatch))
			r.Patch("/{id}/start", endpoints.HandlerError(handler.CampaignStart))
			r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", endpoints.HandlerError(handler.UserLoginPost))
		})
	})

	port := internal.SERVER_PORT
	log.Printf("🚀 Servidor iniciado na porta %s", port)

	if err := http.ListenAndServe(port, routes); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
