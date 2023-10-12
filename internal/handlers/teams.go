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

type TeamHandler struct {
	teamService services.TeamService
}

func NewTeamHandler(teamService services.TeamService) *TeamHandler {
	return &TeamHandler{
		teamService: teamService,
	}
}

func (t *TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var team generated.Team
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(team); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	// TODO: there might be more validation to do here
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	newTeam, err := t.teamService.CreateTeam(r.Context(), &team)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newTeam)
}

func (t *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	queryParams, err := render.ParseQueryFilterParams(r.URL.RawQuery)

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	items, totalFiltered, err := t.teamService.GetTeams(r.Context(), queryParams)

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

func (t *TeamHandler) GetTeamByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	task, err := t.teamService.GetTeamByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "team not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, task)
}

func (t *TeamHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var updateTeam models.UpdateTeamPayload
	var validationErrors []render.ValidationErrorDetails

	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateTeam); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(updateTeam); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	// TODO: there might be more validation to do here
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	defer r.Body.Close()

	updatedTeam, err := t.teamService.UpdateTeam(r.Context(), id, &generated.Team{
		Name:        updateTeam.Name,
		Description: updateTeam.Description,
		Status:      updateTeam.Status,
	})

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "team not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, updatedTeam)
}
