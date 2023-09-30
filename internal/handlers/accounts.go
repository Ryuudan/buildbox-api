package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils/render"
	"golang.org/x/crypto/bcrypt"
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

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var account models.RegisterAccountStruct
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		render.Error(w, r, "json_validation", http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(account); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	user, _ := a.userService.GetUserByEmail(r.Context(), account.Email)

	if user != nil {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "email",
			Message: "email already exists, please try another one",
		})
	}

	plan, err := a.planService.GetPlanByID(r.Context(), account.PlanID)

	if err != nil {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "plan_id",
			Message: err.Error(),
		})
	}

	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	if strings.ToLower(plan.Name) == "demo" {
		account.BillingCycle = "monthly"
		// TODO: make sure to change the discount to 0 when it is a demo account
		// TODO: make sure to make the accounts settings based on demo account
	}

	newAccount, err := a.accountService.CreateAccount(r.Context(), &generated.Account{
		Name:        account.Name,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
	})

	if err != nil {
		render.Error(w, r, "accounts", http.StatusInternalServerError, err.Error())
		return
	}

	_, err = a.subscriptionService.CreateSubscription(r.Context(), &generated.Subscription{
		PlanID:       plan.ID,
		AccountID:    newAccount.ID,
		Status:       "active",
		StartDate:    time.Now(),
		EndDate:      CalculateEndDate(plan.Name, account.BillingCycle),
		BillingCycle: account.BillingCycle,
		Discount:     0, // TODO: make sure this is based on right conditions like when its yearly or any other conditions
	})

	if err != nil {
		render.Error(w, r, "subscription", http.StatusInternalServerError, err.Error())
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Error(w, r, "password", http.StatusBadRequest, err.Error())
		return
	}

	// Creater user for the account
	_, err = a.userService.RegisterUser(r.Context(), &generated.User{
		AccountID:  newAccount.ID,
		Email:      account.Email,
		Password:   string(password),
		FirstName:  account.FirstName,
		LastName:   account.LastName,
		MiddleName: &account.MiddleName,
	})

	if err != nil {
		render.Error(w, r, "accounts", http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	render.JSON(w, http.StatusCreated, newAccount)
}

func (a *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.accountService.GetAccounts(r.Context())
	if err != nil {
		render.Error(w, r, "accounts", http.StatusInternalServerError, err.Error())
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
