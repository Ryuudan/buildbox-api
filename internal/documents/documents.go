package documents

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
// Define package-level constants or configuration settings here.
)

func Initialize(client *mongo.Client) {
	routes := chi.NewRouter()
	println("-------------- DOCUMENTS MODULE INITIALIZED --------------")
	routes.Get("/documents", func(w http.ResponseWriter, r *http.Request) {
		// Write a string response to the client.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a simple endpoint for listing documents"))
	})
}
