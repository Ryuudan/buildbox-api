package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type MilestoneHandler struct {
	milestoneService services.MilestoneService
	taskService      services.TaskService
	projectService   services.ProjectService
}

func NewMilestoneHandler(milestoneService services.MilestoneService, taskService services.TaskService, projectService services.ProjectService) *MilestoneHandler {
	return &MilestoneHandler{
		milestoneService: milestoneService,
		taskService:      taskService,
		projectService:   projectService,
	}
}

func (m *MilestoneHandler) CreateMilestone(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var milestone generated.Milestone
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&milestone); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	defer r.Body.Close()

	// Struct level validation
	if err := validate.Struct(milestone); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	existingProject, _ := m.projectService.GetProjectByID(r.Context(), milestone.ProjectID)

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

	newMilestone, err := m.milestoneService.CreateMilestone(r.Context(), &milestone)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newMilestone)
}

func (m *MilestoneHandler) GetMilestoneByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	project, err := m.milestoneService.GetMilestoneByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "milestone not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, project)
}

func (m *MilestoneHandler) GetMilestones(w http.ResponseWriter, r *http.Request) {

	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	milestones, totalFiltered, err := m.milestoneService.GetMilestones(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: milestones,
		Meta: render.GenerateMeta(
			totalFiltered,
			queryParams,
			len(milestones),
		),
	})
}
