package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func V1Issues(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /issues")

	issuesService := services.NewIssueService(client.Issue)
	projectService := services.NewProjectService(client.Project)
	handler := handlers.NewIssueHandlers(issuesService, projectService)

	router.Route("/issues", func(r chi.Router) {
		r.Get("/", handler.GetIssues)
		r.Post("/", handler.CreateIssue)
		r.Get("/{id}", handler.GetIssueByID)
	})
}
