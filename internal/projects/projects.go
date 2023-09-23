package projects

import (
	"log"

	models "github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/accounts"
	"github.com/go-chi/chi/v5"
)

func Initialize(client *models.Client, router chi.Router) {
	log.Printf("PROJECTS MODULE INITIALIZED")

	projectService := NewProjectService(client)
	accountService := accounts.NewAccountService(client)

	project := NewProjectHandler(projectService, accountService)

	router.Get("/projects", project.GetProjects)
	router.Post("/projects", project.CreateProject)
	router.Get("/projects/{id}", project.GetProjectByID)
}
