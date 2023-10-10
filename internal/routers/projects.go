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
	serviceProviderService := services.NewServiceProviderService(client.ServiceProvider)

	// projectServiceProviderService := services.NewProjectServiceProviderService(client.ProjectServiceProvider)
	project := handlers.NewProjectHandler(projectService, accountService, serviceProviderService)

	router.Route("/projects", func(r chi.Router) {
		r.Get("/", project.GetProjects)
		r.Post("/", project.CreateProject)
		r.Get("/{id}", project.GetProjectByID)
		// r.Get("/{id}/service-providers", project.GetProjectServiceProviders)
		// r.Post("/{id}/service-providers", project.AddServiceProviderToProject)
		// other approach
		// r.Route("/{id}", func(r chi.Router) {
		// 	r.Get("/", project.GetProjectByID)
		// 	// r.Patch("/delete", project.DeleteProjectByID)
		// })
	})
}
