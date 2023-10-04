package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication

func V1Projects(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /projects")
	projectService := services.NewProjectService(client.Project)
	accountService := services.NewAccountService(client.Account)
	project := handlers.NewProjectHandler(projectService, accountService)

	router.Route("/projects", func(r chi.Router) {
		r.Get("/", project.GetProjects)
		r.Post("/", project.CreateProject)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", project.GetProjectByID)
			// r.Patch("/delete", project.DeleteProjectByID)
		})
	})
}
