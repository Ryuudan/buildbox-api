package routers

import (
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

// These routes needs authentication

func V1Projects(client *generated.Client, router chi.Router) {

	projectService := services.NewProjectService(client)
	accountService := services.NewAccountService(client)
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
