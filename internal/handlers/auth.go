package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	accountService      services.AccountService
	userService         services.UserService
	planService         services.PlanService
	subscriptionService services.SubscriptionService
	roleService         services.RolesService
}

func NewAuthHandler(accountService services.AccountService, userService services.UserService, planService services.PlanService, subscriptionService services.SubscriptionService, roleService services.RolesService) *AuthHandler {
	return &AuthHandler{
		accountService:      accountService,
		userService:         userService,
		planService:         planService,
		subscriptionService: subscriptionService,
		roleService:         roleService,
	}
}

func (a *AuthHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var account models.RegisterAccountStruct
	var validationErrors []render.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	defer r.Body.Close()
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
		render.Error(w, r, http.StatusInternalServerError, err.Error())
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
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	role, err := a.roleService.CreateRole(r.Context(), &generated.Role{
		Name:        "Super Admin",
		Description: "Super Admin of the " + account.Name + " account",
		AccountID:   newAccount.ID,
		Permissions: utils.DEFAULT_OWNER_PERMISSIONS,
	})

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
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
		RoleID:     role.ID,
	})

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusCreated, newAccount)
}

func (u *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var credentials models.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		render.Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(credentials); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	user, err := u.userService.GetUserByEmail(r.Context(), credentials.Email)

	if err != nil {
		render.CustomValidationError(w, r, []render.ValidationErrorDetails{
			{
				Field:   "email",
				Message: "Email does not exist.",
			},
		})
		return
	}

	// Compare the provided password with the stored password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

	if err != nil {
		render.CustomValidationError(w, r, []render.ValidationErrorDetails{
			{
				Field:   "password",
				Message: "Invalid password, please try again.",
			},
		})
		return
	}

	subscription, err := u.subscriptionService.GetActiveSubscriptionByAccountID(r.Context(), user.AccountID)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, "Failed to get subscription.")
		return
	}

	if time.Now().After(subscription.EndDate) {
		render.Error(w, r, http.StatusUnauthorized, "Your subscription has expired, please renew your subscription.")
		return
	}

	// TODO: Add Plan and Subscription data to the generated token
	// so that we can directly access that using clains
	accessToken, err := generateAccessToken(user, subscription)
	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, "Failed to generate access token.")
		return
	}

	render.JSON(w, http.StatusOK, map[string]string{"access_token": accessToken})
}

func generateAccessToken(user *generated.User, subscription *generated.Subscription) (string, error) {
	claims := models.CustomClaims{
		AccountID:      user.AccountID,
		PlanID:         subscription.Edges.Plan.ID,
		SubscriptionID: subscription.ID,
		RoleID:         user.RoleID,
		UserID:         user.ID,
		UserUUID:       user.UUID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Subscription: generated.Subscription{
			ID:           subscription.ID,
			StartDate:    subscription.StartDate,
			EndDate:      subscription.EndDate,
			Status:       subscription.Status,
			BillingCycle: subscription.BillingCycle,
			Discount:     subscription.Discount,
		},
		Plan: generated.Plan{
			ID:          subscription.Edges.Plan.ID,
			Name:        subscription.Edges.Plan.Name,
			Description: subscription.Edges.Plan.Description,
			Price:       subscription.Edges.Plan.Price,
		},
		Role:    *user.Edges.Role,
		Account: *user.Edges.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Example: 24 hours from now
		},
	}

	// Set expiration time for the token
	// Generate the JWT token with the payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
