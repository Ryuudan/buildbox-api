package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication
func V1Accounts(client *generated.Client, router chi.Router) {

	accountService := services.NewAccountService(client)
	userService := services.NewUserService(client)

	account := handlers.NewAccountHandler(accountService, userService)

	router.Route("/accounts", func(r chi.Router) {
		// middles wares here
		r.Get("/", account.GetAccounts)
		r.Post("/", account.CreateAccount)
	})

}
