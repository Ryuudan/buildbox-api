package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type ServiceProviderHandler struct {
	serviceProviderService services.ServiceProviderService
}

func NewServiceProviderHandler(serviceProviderService services.ServiceProviderService) *ServiceProviderHandler {
	return &ServiceProviderHandler{
		serviceProviderService: serviceProviderService,
	}
}

func (s *ServiceProviderHandler) GetServiceProviders(w http.ResponseWriter, r *http.Request) {
	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	items, totalFiltered, err := s.serviceProviderService.GetServiceProviders(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: items,
		Meta: render.GenerateMeta(
			totalFiltered,
			queryParams,
			len(items),
		),
	})
}

func (s *ServiceProviderHandler) GetServiceProviderByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	task, err := s.serviceProviderService.GetServiceProviderByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "service provider not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, task)
}

func (s *ServiceProviderHandler) CreateServiceProvider(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()
	var wg sync.WaitGroup
	var mu sync.Mutex

	var service_provider generated.ServiceProvider
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&service_provider); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(service_provider); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	wg.Add(2)
	// TODO: make a function that checks all tables for existing email and phone number to avoid code duplication
	// and to really make sure that no other phone and email exists in another tables

	go func() {
		defer wg.Done()
		existingPhoneNumber, _ := s.serviceProviderService.GetServiceProviderByPhone(r.Context(), service_provider.PhoneNumber)
		if existingPhoneNumber != nil {
			mu.Lock()
			defer mu.Unlock()
			validationErrors = append(validationErrors, render.ValidationErrorDetails{
				Field:   "phone_number",
				Message: "phone number already exists, please try another one",
			})
		}
	}()

	go func() {
		defer wg.Done()
		existingEmail, _ := s.serviceProviderService.GetServiceProviderByEmail(r.Context(), service_provider.Email)
		if existingEmail != nil {
			mu.Lock()
			defer mu.Unlock()
			validationErrors = append(validationErrors, render.ValidationErrorDetails{
				Field:   "email",
				Message: "email already exists, please try another one",
			})
		}
	}()

	wg.Wait()

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	newServiceProvider, err := s.serviceProviderService.CreateServiceProvider(r.Context(), &service_provider)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newServiceProvider)
}
