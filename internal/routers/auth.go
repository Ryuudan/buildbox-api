package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// Public auth means no auth required, this is for login, register, forgot password
// and other public routes
// TODO: make a separate route for Creation of Demo accounts
func V1Public(client *generated.Client, router chi.Router) {

	accountService := services.NewAccountService(client.Account)
	userService := services.NewUserService(client.User)
	subscriptionService := services.NewSubscriptionService(client.Subscription)
	planService := services.NewPlanService(client.Plan)
	rolesService := services.NewRolesService(client.Role)

	authHandler := handlers.NewAuthHandler(accountService, userService, planService, subscriptionService, rolesService)

	router.Post("/register", authHandler.CreateAccount)
	router.Post("/login", authHandler.LoginUser)
}
