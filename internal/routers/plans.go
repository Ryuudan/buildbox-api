package routers

import (
	"github.com/Pyakz/buildbox-api/database"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

// These routes needs authentication
func V1Plans(client *generated.Client, router chi.Router) {
	planService := services.NewPlanService(client.Plan)
	planHandler := handlers.NewPlanHandler(planService, database.Cache(nil))

	router.Route("/plans", func(r chi.Router) {
		r.Get("/{id}", planHandler.GetPlan)
		r.Patch("/{id}", planHandler.UpdatePlan)
	})

}

func V1PublicPlans(client *generated.Client, redis_client *redis.Client, router chi.Router) {

	cache := database.Cache(redis_client)

	planService := services.NewPlanService(client.Plan)
	planHandler := handlers.NewPlanHandler(planService, cache)

	router.Route("/plans", func(r chi.Router) {
		r.Get("/", planHandler.GetPlans)
		r.Post("/", planHandler.CreatePlan)

	})
}
