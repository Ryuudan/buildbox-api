package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Users(client *generated.Client, router chi.Router) {

	userService := services.NewUserService(client.User)

	user := handlers.NewUserHandler(userService)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", user.CreateUser)
	})

}
