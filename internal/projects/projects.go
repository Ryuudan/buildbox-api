package projects

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
)

func Initialize(client *ent.ProjectClient, router chi.Router) {
	log.Printf("PROJECTS MODULE INITIALIZED")
	projectService := NewProjectService(client)
	project := NewProjectHandler(projectService)

	router.Route("/v1", func(v1 chi.Router) {
		v1.Use(utils.VersionMiddleware("v1"))
		v1.Get("/projects", project.GetProjects)
		v1.Post("/projects", project.CreateProject)
		v1.Get("/projects/{id}", project.GetProjectByID)

	})
}
