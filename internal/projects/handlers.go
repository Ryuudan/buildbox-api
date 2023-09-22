package projects

import (
	"encoding/json"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent"
	"github.com/Pyakz/buildbox-api/utils"
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
		utils.RenderError(w, "Invalid JSON format. Please check your request and try again.", http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(project); err != nil {
		utils.Validate(w, err)
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
		utils.CustomValidationError(w, validationErrors)
		return
	}

	newProject, err := p.projectService.CreateProject(r.Context(), &project)

	if err != nil {
		utils.RenderError(w, "Oops! Something went wrong while creating the project. Please try again later.", http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	utils.RenderJSON(w, http.StatusCreated, newProject)
}

func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.projectService.GetProjects(r.Context())
	if err != nil {
		// Handle the error, e.g., by sending an error response.
		http.Error(w, "Failed to retrieve projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RenderJSON(w, http.StatusOK, projects)
}

func (p *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, http.StatusOK, []string{"Project 1", "Project 2", "Project 3"})

}
