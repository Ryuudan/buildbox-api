package handlers

import (
	"github.com/Pyakz/buildbox-api/internal/services"
)

type SearchHandler struct {
	accountService      services.AccountService
	userService         services.UserService
	planService         services.PlanService
	subscriptionService services.SubscriptionService
	roleService         services.RolesService
}

func NewSearchHandler(accountService services.AccountService, userService services.UserService, planService services.PlanService, subscriptionService services.SubscriptionService, roleService services.RolesService) *SearchHandler {
	return &SearchHandler{
		accountService:      accountService,
		userService:         userService,
		planService:         planService,
		subscriptionService: subscriptionService,
		roleService:         roleService,
	}
}

// TODO: Search Endpoing
// func (s *SearchHandler) GlobalSearch(w http.ResponseWriter, r *http.Request) {
//
// }
