package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/middlewares"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func PrivateInitializeRoutes(client *generated.Client, router chi.Router) http.Handler {
	log.Println("✅ Private Routes Initialized")

	validity := middlewares.NewUserMiddleware(services.NewUserService(client.User))

	private := chi.NewRouter()

	private.Use(utils.VersionMiddleware("1.0"))
	private.Use(httprate.LimitByIP(300, 1*time.Minute))
	private.Use(middlewares.AuthMiddleware)
	private.Use(validity.ValidUser)

	private.Group(func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("1.0"))
		V1Accounts(client, v1)
		V1Projects(client, v1)
		V1Plans(client, v1)
	})

	router.Mount("/api/v1", private)
	return router
}

func PublicInitializeRoutes(client *generated.Client, router chi.Router) http.Handler {
	log.Println("✅ Public Routes Initialized")

	public := chi.NewRouter()

	public.Use(httprate.LimitByIP(20, 1*time.Minute))
	public.Use(middleware.Heartbeat("/"))

	public.Group(func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("1.0"))
		V1Public(client, v1)
		V1PublicPlans(client, v1)
	})

	router.Mount("/public/v1", public)
	return router
}
