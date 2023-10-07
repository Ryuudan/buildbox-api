package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Users(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /users")
	userService := services.NewUserService(client.User)
	subscriptionService := services.NewSubscriptionService(client.Subscription)

	user := handlers.NewUserHandler(userService, subscriptionService)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", user.CreateUser)
		r.Get("/", user.GetUsers)
		r.Get("/{id}", user.GetUserByID)
	})
}
