package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils/render"
)

type TaskHandlers struct {
	taskService    services.TaskService
	projectService services.ProjectService
}

func NewTaskHandlers(taskService services.TaskService, projectService services.ProjectService) *TaskHandlers {
	return &TaskHandlers{
		taskService:    taskService,
		projectService: projectService,
	}
}

func (t *TaskHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	var filters models.Filters

	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	orders, err := render.ParseOrderString(queryParams.Order)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	for _, fields := range orders {
		filters.Order = append(filters.Order, *fields)
	}

	tasks, total, err := t.taskService.GetTasks(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: tasks,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(tasks),
		),
	})
}

func (t *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var task generated.Task
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(task); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	existingProject, _ := t.projectService.GetProjectByID(r.Context(), task.ProjectID)

	if existingProject == nil {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "project_id",
			Message: "Project does not exist",
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	newTask, err := t.taskService.CreateTask(r.Context(), &task)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	render.JSON(w, http.StatusCreated, newTask)
}
