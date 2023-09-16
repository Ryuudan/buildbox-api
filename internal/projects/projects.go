package projects

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
// Define package-level constants or configuration settings here.
)

func Initialize(client *mongo.Client, router chi.Router) {
	println("-------------- PROJECTS MODULE INITIALIZED --------------")
	router.Get("/projects", func(w http.ResponseWriter, _ *http.Request) {
		// Write a string response to the client.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a simple endpoint for listing projects"))
	})
}
