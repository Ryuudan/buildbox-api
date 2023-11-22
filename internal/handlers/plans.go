package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/database"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type PlanHandler struct {
	planService services.PlanService
	cache       *database.RedisCache
}

func NewPlanHandler(planService services.PlanService, cache *database.RedisCache) *PlanHandler {
	return &PlanHandler{
		planService: planService,
		cache:       cache,
	}
}

func (p *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {

	validate := render.Validator()

	var plan generated.Plan
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}
	defer r.Body.Close()
	// Struct level validation
	if err := validate.Struct(plan); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	existingName, _ := p.planService.GetPlanByName(r.Context(), plan.Name)
	if existingName != nil {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "name",
			Message: fmt.Sprintf("Plan with name '%s' already exists", plan.Name),
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	plan.Name = strings.ToLower(plan.Name)
	newPlan, err := p.planService.CreatePlan(r.Context(), &plan)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// clear the cache
	p.cache.ClearCache("plans")
	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newPlan)
}

func (p *PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve plans from cache
	if existing, err := p.cache.GetCache("plans"); err == nil && existing != "" {
		var plans []generated.Plan
		if err := json.Unmarshal([]byte(existing), &plans); err == nil {
			render.JSON(w, http.StatusOK, plans)
			return
		}
	}

	// Plans not found in the cache or error occurred, fetch from the service
	plans, err := p.planService.GetPlans(r.Context())
	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// If plans are obtained, store them in the cache
	if len(plans) > 0 {
		if plansJSON, err := json.Marshal(plans); err == nil {
			if err := p.cache.SetCache("plans", string(plansJSON), 0); err != nil {
				fmt.Println("Error setting cache:", err)
			}
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Return the plans
	render.JSON(w, http.StatusOK, plans)
}

func (p *PlanHandler) GetPlan(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	plan, err := p.planService.GetPlanByID(r.Context(), id)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "plan not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, plan)
}

func (p *PlanHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	var updated models.UpdatePlanRequest
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	defer r.Body.Close()

	// Struct level validation
	if err := validate.Struct(updated); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	*updated.Name = strings.ToLower(*updated.Name)

	existingPlan, _ := p.planService.GetPlanByName(r.Context(), *updated.Name)

	if existingPlan != nil && existingPlan.Name == *updated.Name && existingPlan.ID != id {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "name",
			Message: fmt.Sprintf("Plan with name '%s' already exists", *updated.Name),
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	plan, err := p.planService.UpdatePlan(r.Context(), id, updated)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "plan not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, plan)
}
