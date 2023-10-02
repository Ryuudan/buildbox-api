package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils/render"
)

type PlanHandler struct {
	planService services.PlanService
}

func NewPlanHandler(planService services.PlanService) *PlanHandler {
	return &PlanHandler{
		planService: planService,
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

	newPlan, err := p.planService.CreatePlan(r.Context(), &plan)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	render.JSON(w, http.StatusCreated, newPlan)
}

func (p *PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := p.planService.GetPlans(r.Context())

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, plans)
}
