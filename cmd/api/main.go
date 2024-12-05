package main

import (
	"log"
	"net/http"
	"senderEmails/internal/domain/campaign"
	"senderEmails/internal/endpoints"
	"senderEmails/internal/infraestructure/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	routes := chi.NewRouter()

	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	db := database.NewConnectionToDB()

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{
			Db: db,
		},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	routes.Route("/api/v1", func(r chi.Router) {
		r.Route("/campaigns", func(r chi.Router) {
			r.Post("/", endpoints.HandlerError(handler.CampaignPost))
			r.Get("/", endpoints.HandlerError(handler.CampaignsGet))
			r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
			r.Patch("/{id}/cancel", endpoints.HandlerError(handler.CampaignCancelPatch))
			r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		})
	})

	port := ":5000"
	log.Printf("🚀 Servidor iniciado na porta %s", port)

	if err := http.ListenAndServe(port, routes); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
