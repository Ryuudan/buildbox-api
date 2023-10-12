package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type TaskHandlers struct {
	taskService      services.TaskService
	projectService   services.ProjectService
	milestoneService services.MilestoneService
}

func NewTaskHandlers(taskService services.TaskService, projectService services.ProjectService, milestoneService services.MilestoneService) *TaskHandlers {
	return &TaskHandlers{
		taskService:      taskService,
		projectService:   projectService,
		milestoneService: milestoneService,
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

	if task.TaskMilestoneID != nil {
		existingMilestone, _ := t.milestoneService.GetMilestoneByID(r.Context(), *task.TaskMilestoneID)

		if existingMilestone == nil {
			validationErrors = append(validationErrors, render.ValidationErrorDetails{
				Field:   "task_milestone_id",
				Message: "Milestone does not exist",
			})
		}
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

func (t *TaskHandlers) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	task, err := t.taskService.GetTaskByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "task not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}
	render.JSON(w, http.StatusOK, task)
}
