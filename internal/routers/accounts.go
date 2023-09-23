package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func Accounts(client *generated.Client, router chi.Router) {
	projectService := services.NewProjectService(client)
	accountService := services.NewAccountService(client)

	account := handlers.NewAccountHandler(accountService, projectService)

	router.Get("/accounts", account.GetAccounts)
	router.Post("/accounts", account.CreateAccount)
	log.Println("✅ Accounts Module Initialized")

}
