package projects

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	projectService ProjectService
}

func NewProjectHandler(projectService ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (p *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	validate := utils.Validator()

	var project ent.Project
	var validationErrors []utils.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.RenderError(w, r, "projects", http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(project); err != nil {
		utils.Validate(w, r, err)
		return
	}

	// this is sample validation error we can call on our own
	// if somethingHappend  {
	// 	validationErrors = append(validationErrors, utils.ValidationErrorDetails{
	// 		Field:   "somthing",
	// 		Message: "something",
	// 	})
	// }

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		utils.CustomValidationError(w, r, validationErrors)
		return
	}

	newProject, err := p.projectService.CreateProject(r.Context(), &project)

	if err != nil {
		utils.RenderError(w, r, "projects", http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	utils.RenderJSON(w, http.StatusCreated, newProject)
}

func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.projectService.GetProjects(r.Context())
	if err != nil {
		utils.RenderError(w, r, "projects", http.StatusInternalServerError, err.Error())
		return
	}

	utils.RenderJSON(w, http.StatusOK, projects)
}

func (p *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		utils.RenderError(w, r, "projects", http.StatusBadRequest, "Invalid ID")
		return
	}

	project, err := p.projectService.GetProjectByID(r.Context(), id)
	if err != nil {
		utils.RenderError(w, r, "projects", http.StatusBadRequest, err.Error())
		return
	}

	utils.RenderJSON(w, http.StatusOK, project)
}
