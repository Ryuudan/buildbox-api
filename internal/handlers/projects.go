package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	projectService services.ProjectService
	accountService services.AccountService
}

func NewProjectHandler(projectService services.ProjectService, accountService services.AccountService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		accountService: accountService,
	}
}

func (p *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var project generated.Project
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		render.Error(w, r, "json_validation", http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(project); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	newProject, err := p.projectService.CreateProject(r.Context(), &project)

	if err != nil {
		render.Error(w, r, "projects", http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	render.JSON(w, http.StatusCreated, newProject)
}

func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {

	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, "Params Error", http.StatusBadRequest, err.Error())
		return
	}

	projects, total, err := p.projectService.GetProjects(r.Context(), queryParams)
	if err != nil {
		render.Error(w, r, "projects", http.StatusInternalServerError, err.Error())
		return
	}

	response := &render.PaginatedResults{
		Results: projects,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(projects),
		),
	}

	render.JSON(w, http.StatusOK, response)
}

func (p *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, "projects", http.StatusBadRequest, "Invalid ID")
		return
	}

	project, err := p.projectService.GetProjectByID(r.Context(), id)
	if err != nil {
		render.Error(w, r, "projects", http.StatusBadRequest, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, project)
}
