package accounts

import (
	"log"

	models "github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/go-chi/chi/v5"
)

func Initialize(client *models.Client, router chi.Router) {
	log.Printf("ACCOUNTS MODULE INITIALIZED")
	accountService := NewAccountService(client)
	account := NewAccountHandler(accountService)

	router.Get("/accounts", account.GetAccounts)
	router.Post("/accounts", account.CreateAccount)
}
