package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Accounts(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /accounts")
	accountService := services.NewAccountService(client.Account)
	userService := services.NewUserService(client.User)
	planService := services.NewPlanService(client.Plan)
	subscriptionService := services.NewSubscriptionService(client.Subscription)

	account := handlers.NewAccountHandler(accountService, userService, planService, subscriptionService)

	router.Route("/accounts", func(r chi.Router) {
		r.Get("/", account.GetAccounts)
	})

}
