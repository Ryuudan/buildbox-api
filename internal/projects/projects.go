package projects

import (
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Initialize(mongo *mongo.Database, router chi.Router) {
	println("-------------- PROJECTS MODULE INITIALIZED --------------")

	service := NewProjectService(mongo.Collection("projects"))
	project := NewProjectHandler(service)

	router.Route("/v1", func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("v1"))

		v1.Get("/projects", project.GetProjectsHandler)
		v1.Get("/projects/{id}", project.GetProjectHandler)
		v1.Post("/projects", project.CreateProjectHandler)

	})
}
