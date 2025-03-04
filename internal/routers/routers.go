package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/middlewares"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/redis/go-redis/v9"
)

func PrivateInitializeRoutes(client *generated.Client, router chi.Router) http.Handler {
	private := chi.NewRouter()

	private.Use(httprate.LimitByIP(300, 1*time.Minute))
	private.Use(middlewares.AuthMiddleware)

	private.Group(func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("1.0"))
		V1Accounts(client, v1)
		V1Projects(client, v1)
		V1Users(client, v1)
		V1Plans(client, v1)
		V1Roles(client, v1)
		V1Tasks(client, v1)
		V1Milestones(client, v1)
		V1Issues(client, v1)
		V1ServiceProviders(client, v1)
		V1Teams(client, v1)
		// TODO: Search route
	})

	router.Mount("/api/v1", private)
	return router
}

func PublicInitializeRoutes(client *generated.Client, redis_client *redis.Client, router chi.Router) http.Handler {
	log.Println("✅ Routes: /public")
	public := chi.NewRouter()

	public.Use(httprate.LimitByIP(20, 1*time.Minute))

	public.Group(func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("1.0"))
		V1Public(client, v1)
		V1PublicPlans(client, redis_client, v1)
	})

	router.Mount("/public/v1", public)
	return router
}
