package projects

import (
	"net/http"

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
func (p *ProjectHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, http.StatusOK, []string{"Project 1", "Project 2", "Project 3"})

}

// func (p *ProjectHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	validate := validator.New()
// 	var project Project
// 	var validationErrors []utils.ValidationErrorDetails

// 	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
// 		utils.RenderError(w, "Invalid JSON format. Please check your request and try again.", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	// Struct level validation
// 	if err := validate.Struct(project); err != nil {
// 		utils.Validate(w, err)
// 		return
// 	}

// 	if project.StartDate != nil && project.EndDate != nil && project.StartDate.After(*project.EndDate) {
// 		validationErrors = append(validationErrors, utils.ValidationErrorDetails{
// 			Field:   "startDate",
// 			Message: "Start date must be before end date",
// 		})
// 	}

// 	if project.Budget < 0 {
// 		validationErrors = append(validationErrors, utils.ValidationErrorDetails{
// 			Field:   "budget",
// 			Message: "Budget must be greater than or equal to 0",
// 		})
// 	}

// 	// If there are validation errors, return a custom validation error response
// 	if len(validationErrors) > 0 {
// 		utils.CustomValidationError(w, validationErrors)
// 		return
// 	}

// 	newProject, err := p.projectService.CreateProject(r.Context(), &project)

// 	if err != nil {
// 		utils.RenderError(w, "Oops! Something went wrong while creating the project. Please try again later.", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	defer r.Body.Close()
// 	utils.RenderJSON(w, http.StatusCreated, newProject)
// }

func (p *ProjectHandler) GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// You can call methods from ph.projectService to fetch projects from your service
	// For example:
	// projects, err := ph.projectService.GetProjects(r.Context(), queryParams, filters)
	// Handle the result and error accordingly
	utils.RenderJSON(w, http.StatusOK, []string{"Project 1", "Project 2", "Project 3"})
}

func (p *ProjectHandler) GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, http.StatusOK, []string{"Project 1", "Project 2", "Project 3"})

}
