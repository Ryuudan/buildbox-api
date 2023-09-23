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

func InitializeRoutes(client *generated.Client, router chi.Router) http.Handler {
	log.Println("✅ Routes Initialized")

	validity := middlewares.NewUserMiddleware(services.NewUserService(client))

	v1 := chi.NewRouter()
	v1.Use(utils.VersionMiddleware("1.0"))

	v1.Group(func(private chi.Router) {
		private.Use(httprate.LimitByIP(300, 1*time.Minute))
		private.Use(middlewares.AuthMiddleware)
		private.Use(validity.ValidUser)

		V1Accounts(client, private)
		V1Projects(client, private)
	})

	v1.Group(func(public chi.Router) {
		public.Use(httprate.LimitByIP(20, 1*time.Minute))
		public.Use(middleware.Heartbeat("/"))
		V1Public(client, public)
	})

	router.Mount("/api/v1", v1)

	return router
}
