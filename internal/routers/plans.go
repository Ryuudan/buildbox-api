package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Plans(client *generated.Client, router chi.Router) {
	planService := services.NewPlanService(client.Plan)
	planHandler := handlers.NewPlanHandler(planService)

	router.Route("/plans", func(r chi.Router) {
		r.Get("/{id}", planHandler.GetPlan)
		r.Patch("/{id}", planHandler.UpdatePlan)
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
