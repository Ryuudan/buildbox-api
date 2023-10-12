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
	log.Println("✅ Routes: /roles")

	rolesService := services.NewRolesService(client.Role)
	handler := handlers.NewRolesHandlers(rolesService)

	router.Route("/roles", func(r chi.Router) {
		r.Post("/", handler.CreateRole)
		r.Get("/", handler.GetRoles)
		r.Patch("/{id}", handler.UpdateRole)
	})
}
