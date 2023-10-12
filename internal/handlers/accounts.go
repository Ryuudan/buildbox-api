package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Pyakz/buildbox-api/constants"
	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
)

type AccountHandler struct {
	accountService      services.AccountService
	userService         services.UserService
	planService         services.PlanService
	subscriptionService services.SubscriptionService
}

// TODO: make a separate fuction for Creation of Demo accounts
func NewAccountHandler(accountService services.AccountService, userService services.UserService, planService services.PlanService, subscriptionService services.SubscriptionService) *AccountHandler {
	return &AccountHandler{
		accountService:      accountService,
		userService:         userService,
		planService:         planService,
		subscriptionService: subscriptionService,
	}
}

func (a *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {

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

	accounts, total, err := a.accountService.GetAccounts(r.Context(), queryParams)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: accounts,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(accounts),
		),
	})
}

func (a *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	account, err := a.accountService.GetAccountByID(r.Context(), id)
	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "Account not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	render.JSON(w, http.StatusOK, account)
}

func CalculateEndDate(planName string, billingCycle subscription.BillingCycle) time.Time {
	var duration time.Duration

	switch billingCycle {
	case "monthly":
		duration = time.Hour * 24 * 30
	case "yearly":
		duration = time.Hour * 24 * 365
	default:
		duration = time.Hour * 24 * 30 // Default to 30 days for unsupported cases
	}

	if strings.ToLower(planName) == "demo" {
		duration = time.Hour * 24 * 30
	}

	return time.Now().Add(duration)
}
