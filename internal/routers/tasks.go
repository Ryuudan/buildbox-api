package routers

import (
	"log"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/handlers"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func V1Tasks(client *generated.Client, router chi.Router) {
	log.Println("✅ Routes: /tasks")

	taskService := services.NewTaskService(client.Task)
	projectService := services.NewProjectService(client.Project)
	taskHandler := handlers.NewTaskHandlers(taskService, projectService)

	router.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.GetTasks)
		r.Post("/", taskHandler.CreateTask)
	})
}
