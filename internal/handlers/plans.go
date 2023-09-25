package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
)

type PlanHandler struct {
	planService services.PlanService
}

func NewPlanHandler(planService services.PlanService) *PlanHandler {
	log.Println("✅ Plans Handler Initialized")
	return &PlanHandler{
		planService: planService,
	}
}

func (p *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {

	validate := utils.Validator()

	var plan models.CreatePlanRequest
	var validationErrors []utils.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		utils.RenderError(w, r, "json_validation", http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(plan); err != nil {
		utils.Validate(w, r, err)
		return
	}

	existingName, _ := p.planService.GetPlanByName(r.Context(), plan.Name)
	if existingName != nil {
		validationErrors = append(validationErrors, utils.ValidationErrorDetails{
			Field:   "name",
			Message: fmt.Sprintf("Plan with name '%s' already exists", plan.Name),
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		utils.CustomValidationError(w, r, validationErrors)
		return
	}

	newPlan, err := p.planService.CreatePlan(r.Context(), &generated.Plan{
		Name:        plan.Name,
		Description: plan.Description,
		Price:       &plan.Price,
	})

	if err != nil {
		utils.RenderError(w, r, "plans", http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	utils.RenderJSON(w, http.StatusCreated, newPlan)
}

func (p *PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := p.planService.GetPlans(r.Context())

	if err != nil {
		utils.RenderError(w, r, "plans", http.StatusInternalServerError, err.Error())
		return
	}

	utils.RenderJSON(w, http.StatusOK, plans)
}
