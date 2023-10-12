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

type IssueHandlers struct {
	issueService   services.IssueService
	projectService services.ProjectService
}

func NewIssueHandlers(issueService services.IssueService, projectService services.ProjectService) *IssueHandlers {
	return &IssueHandlers{
		issueService:   issueService,
		projectService: projectService,
	}
}

func (i *IssueHandlers) CreateIssue(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var issue generated.Issue
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(issue); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	existingProject, _ := i.projectService.GetProjectByID(r.Context(), issue.ProjectID)

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

	newIssue, err := i.issueService.CreateIssue(r.Context(), &issue)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	render.JSON(w, http.StatusCreated, newIssue)
}

func (i *IssueHandlers) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	project, err := i.issueService.GetIssueByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "issue not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, project)
}

func (i *IssueHandlers) GetIssues(w http.ResponseWriter, r *http.Request) {

	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	issues, totalFiltered, err := i.issueService.GetIssues(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: issues,
		Meta: render.GenerateMeta(
			totalFiltered,
			queryParams,
			len(issues),
		),
	})
}
