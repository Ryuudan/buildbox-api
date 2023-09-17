package materials

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
// Define package-level constants or configuration settings here.
)

func Initialize(db *mongo.Database, router *chi.Mux) {
	println("-------------- MATERIALS MODULE INITIALIZED --------------")
	router.Get("/materials", func(w http.ResponseWriter, _ *http.Request) {
		// Write a string response to the client.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a simple endpoint for listing materials"))
	})
}
