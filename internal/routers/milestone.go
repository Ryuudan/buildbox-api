package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func V1Milestones(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /milestones")

	milestoneService := services.NewMilestoneService(client.Milestone)
	taskService := services.NewTaskService(client.Task)
	projectService := services.NewProjectService(client.Project)
	milestoneHandler := handlers.NewMilestoneHandler(milestoneService, taskService, projectService)

	router.Route("/milestones", func(r chi.Router) {
		r.Get("/", milestoneHandler.GetMilestones)
		r.Post("/", milestoneHandler.CreateMilestone)
		r.Get("/{id}", milestoneHandler.GetMilestoneByID)
	})
}
