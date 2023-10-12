package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func V1Teams(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /teams")

	teamService := services.NewTeamService(client.Team)
	handler := handlers.NewTeamHandler(teamService)

	router.Route("/teams", func(r chi.Router) {
		r.Post("/", handler.CreateTeam)
		r.Get("/", handler.GetTeams)
		r.Get("/{id}", handler.GetTeamByID)
		r.Patch("/{id}", handler.UpdateTeam)
	})
}
