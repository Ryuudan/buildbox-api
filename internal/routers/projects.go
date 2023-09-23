package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func Projects(client *generated.Client, router chi.Router) {

	projectService := services.NewProjectService(client)
	accountService := services.NewAccountService(client)

	project := handlers.NewProjectHandler(projectService, accountService)

	router.Get("/projects", project.GetProjects)
	router.Post("/projects", project.CreateProject)
	router.Get("/projects/{id}", project.GetProjectByID)

	log.Println("✅ Projects Module Initialized")

}
