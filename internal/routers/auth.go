package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// Public auth means no auth required, this is for login, register, forgot password
// and other public routes
func V1Public(client *generated.Client, router chi.Router) {

	accountService := services.NewAccountService(client)
	userService := services.NewUserService(client)

	account := handlers.NewAccountHandler(accountService, userService)
	user := handlers.NewUserHandler(userService)

	router.Post("/register", account.CreateAccount)
	router.Post("/login", user.LoginUser)

	// router.Post("/forgot-password", func(w http.ResponseWriter, _ *http.Request) {
	// 	w.Write([]byte("forgot-password"))
	// })

}
