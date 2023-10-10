package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func V1ServiceProviders(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /service-providers")

	serviceProviderService := services.NewServiceProviderService(client.ServiceProvider)
	handler := handlers.NewServiceProviderHandler(serviceProviderService)

	router.Route("/service-providers", func(r chi.Router) {
		r.Post("/", handler.CreateServiceProvider)
		r.Get("/", handler.GetServiceProviders)
		r.Get("/{id}", handler.GetServiceProviderByID)
	})
}
