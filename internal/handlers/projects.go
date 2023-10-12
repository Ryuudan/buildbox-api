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

type ProjectHandler struct {
	projectService         services.ProjectService
	accountService         services.AccountService
	serviceProviderService services.ServiceProviderService
}

func NewProjectHandler(projectService services.ProjectService, accountService services.AccountService, serviceProviderService services.ServiceProviderService) *ProjectHandler {
	return &ProjectHandler{
		projectService:         projectService,
		accountService:         accountService,
		serviceProviderService: serviceProviderService,
	}
}

func (p *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var project generated.Project
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
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
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	render.JSON(w, http.StatusCreated, newProject)
}

func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
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

	projects, total, err := p.projectService.GetProjects(r.Context(), queryParams, filters)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: projects,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(projects),
		),
	})
}

func (p *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	project, err := p.projectService.GetProjectByID(r.Context(), id)
	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "project not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, project)
}

// func (p *ProjectHandler) GetProjectServiceProviders(w http.ResponseWriter, r *http.Request) {

// 	id, err := utils.StringToInt(chi.URLParam(r, "id"))

// 	if err != nil {
// 		render.Error(w, r, http.StatusBadRequest, "Invalid ID")
// 		return
// 	}

// 	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

// 	if err != nil {
// 		render.Error(w, r, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	service_providers, total, err := p.projectServiceProvidersService.GetProjectServiceProviders(
// 		r.Context(),
// 		id,
// 		queryParams,
// 	)

// 	if err != nil {
// 		render.Error(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	render.JSON(w, http.StatusOK, &render.PaginatedResults{
// 		Results: service_providers,
// 		Meta: render.GenerateMeta(
// 			total,
// 			queryParams,
// 			len(service_providers),
// 		),
// 	})
// }

// func (p *ProjectHandler) AddServiceProviderToProject(w http.ResponseWriter, r *http.Request) {

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	id, err := utils.StringToInt(chi.URLParam(r, "id"))

// 	if err != nil {
// 		render.Error(w, r, http.StatusBadRequest, "Invalid ID")
// 		return
// 	}

// 	validate := render.Validator()

// 	var payload models.AddServiceProviderToProjectRequest
// 	var validationErrors []render.ValidationErrorDetails

// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
// 		return
// 	}

// 	// Struct level validation
// 	if err := validate.Struct(payload); err != nil {
// 		render.ValidationError(w, r, err)
// 		return
// 	}

// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()
// 		existingProject, _ := p.projectService.GetProjectByID(r.Context(), id)
// 		if existingProject == nil {
// 			validationErrors = append(validationErrors, render.ValidationErrorDetails{
// 				Field:   "project_id",
// 				Message: "project does not exist",
// 			})
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()
// 		existingProject, _ := p.serviceProviderService.GetServiceProviderByID(r.Context(), payload.ServiceProviderID)
// 		if existingProject == nil {
// 			validationErrors = append(validationErrors, render.ValidationErrorDetails{
// 				Field:   "service_provider_id",
// 				Message: "service provider does not exist",
// 			})
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()
// 		existing, _ := p.projectServiceProvidersService.GetExisting(r.Context(), id, payload.ServiceProviderID)
// 		if existing != nil && existing.ProjectProjectID == id && existing.ProjectServiceProviderID == payload.ServiceProviderID {
// 			validationErrors = append(validationErrors, render.ValidationErrorDetails{
// 				Field:   "service_provider_id",
// 				Message: "service provider already exists in the current project",
// 			})
// 		}
// 	}()

// 	wg.Wait()

// 	// If there are validation errors, return a custom validation error response
// 	if len(validationErrors) > 0 {
// 		render.CustomValidationError(w, r, validationErrors)
// 		return
// 	}

// 	projectServiceProvider, err := p.projectServiceProvidersService.AddServiceProviderToProject(r.Context(), id, payload.ServiceProviderID)

// 	if err != nil {
// 		render.Error(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	defer r.Body.Close()

// 	render.JSON(w, http.StatusCreated, projectServiceProvider)
// }
