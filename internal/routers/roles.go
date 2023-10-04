package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Roles(client *generated.Client, router chi.Router) {
	log.Println("✅ Roles Routes Initialized")

	rolesService := services.NewRolesService(client.Role)
	rolesHandler := handlers.NewRolesHandlers(rolesService)

	router.Route("/roles", func(r chi.Router) {
		r.Post("/", rolesHandler.CreateRole)
		r.Get("/", rolesHandler.GetRoles)
	})
}
