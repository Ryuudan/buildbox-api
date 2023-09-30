package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Plans(client *generated.Client, router chi.Router) {
	log.Println("✅ Plans Routes Initialized")

	planService := services.NewPlanService(client.Plan)
	planHandler := handlers.NewPlanHandler(planService)

	router.Route("/plans", func(r chi.Router) {
		r.Post("/", planHandler.CreatePlan)
	})

}

func V1PublicPlans(client *generated.Client, router chi.Router) {

	planService := services.NewPlanService(client.Plan)
	planHandler := handlers.NewPlanHandler(planService)

	router.Route("/plans", func(r chi.Router) {
		r.Get("/", planHandler.GetPlans)
		r.Post("/", planHandler.CreatePlan)
	})
}
