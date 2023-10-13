package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type RolesHandlers struct {
	rolesService services.RolesService
}

func NewRolesHandlers(rolesService services.RolesService) *RolesHandlers {
	return &RolesHandlers{
		rolesService: rolesService,
	}
}

// TODO: Add permissions
func (ro *RolesHandlers) CreateRole(w http.ResponseWriter, r *http.Request) {

	validate := render.Validator()

	var role generated.Role
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	defer r.Body.Close()
	// Struct level validation
	if err := validate.Struct(role); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	existingName, _ := ro.rolesService.GetRoleByName(r.Context(), role.Name)

	if existingName != nil {
		render.CustomValidationError(w, r, []render.ValidationErrorDetails{
			{
				Field:   "name",
				Message: fmt.Sprintf("Role with name '%s' already exists in this account.", role.Name),
			},
		})
		return
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	newRole, err := ro.rolesService.CreateRole(r.Context(), &role)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newRole)
}

func (ro *RolesHandlers) GetRoles(w http.ResponseWriter, r *http.Request) {
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

	roles, total, err := ro.rolesService.GetRoles(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: roles,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(roles),
		),
	})
}

func (ro *RolesHandlers) UpdateRole(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var payload models.UpdateRolePayload
	var validationErrors []render.ValidationErrorDetails

	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	defer r.Body.Close()

	// Struct level validation
	if err := validate.Struct(payload); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	// TODO: there might be more validation to do here
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	updatedRole, err := ro.rolesService.UpdateRole(r.Context(), id, payload)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "role not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, updatedRole)
}

func (ro *RolesHandlers) GetRole(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	role, err := ro.rolesService.GetRoleByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "role not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, role)
}
