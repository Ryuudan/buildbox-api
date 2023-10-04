package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils/render"
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
	accounts, err := a.accountService.GetAccounts(r.Context())
	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, accounts)
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
