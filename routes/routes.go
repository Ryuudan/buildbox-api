package routes

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/accounts"
	"github.com/Pyakz/buildbox-api/internal/projects"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
)

func Initialize(router chi.Router, client *generated.Client) {
	log.Printf("ROUTES MODULE INITIALIZED")

	// v1 routes
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(utils.VersionMiddleware("v1"))
		projects.Initialize(client, r)
		accounts.Initialize(client, r)
	})

}
