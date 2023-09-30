package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService         services.UserService
	subscriptionService services.SubscriptionService
}

func NewUserHandler(userService services.UserService, subscriptionService services.SubscriptionService) *UserHandler {
	return &UserHandler{
		userService:         userService,
		subscriptionService: subscriptionService,
	}
}

func (u *UserHandler) GetUsersByAccount(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetUsersByAccountID(r.Context(), 1)
	if err != nil {
		render.Error(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}
	render.JSON(w, http.StatusOK, users)
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	user, err := u.userService.GetUserById(r.Context(), 1)
	if err != nil {
		render.Error(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}
	render.JSON(w, http.StatusOK, user)
}

// CreateUser handles the creation of a new user.
// It expects a JSON-encoded user object in the request body,
// validates the input, and then registers the user in the system.
//
// If successful, it returns a JSON response with the newly created user.
//
// If there are validation errors, it returns a custom validation error response.
// If any other errors occur during user creation, it returns an internal server error response.
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Create a validator instance for input validation
	validate := render.Validator()

	var user generated.User
	var validationErrors []render.ValidationErrorDetails

	// Decode the JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.Error(w, r, "json_validation", http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
		return
	}

	// Struct level validation of the user object
	if err := validate.Struct(user); err != nil {
		render.ValidationError(w, r, err)
		return
	}

	// Check if a user with the same email already exists
	existingUser, _ := u.userService.GetUserByEmail(r.Context(), user.Email)

	if existingUser != nil {
		validationErrors = append(validationErrors, render.ValidationErrorDetails{
			Field:   "email",
			Message: "email already exists, please try another one",
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		render.CustomValidationError(w, r, validationErrors)
		return
	}

	// Retrieve user claims from the JWT token in the request context
	claims, ok := r.Context().Value(models.ContextKeyClaims).(jwt.MapClaims)
	if !ok {
		render.Error(w, r, "claims", http.StatusBadRequest, "failed to get user claims from context")
		return
	}

	// Assign the account ID from claims to the user
	user.AccountID = int(claims["account_id"].(float64))

	// Generate a salted and hashed password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Error(w, r, "bcrypt", http.StatusInternalServerError, "failed to generate hashed password")
		return
	}

	// Replace the plaintext password with the hashed password
	user.Password = string(password)

	// Register the user in the system
	newUser, err := u.userService.RegisterUser(r.Context(), &user)

	if err != nil {
		render.Error(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the newly created user in JSON format
	render.JSON(w, http.StatusOK, newUser)
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	validate := render.Validator()

	var credentials models.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		render.Error(w, r, "users", http.StatusBadRequest, err.Error())
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
				Message: "Invalid email or password",
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
		render.Error(w, r, "auth", http.StatusInternalServerError, "Failed to get subscription.")
		return
	}

	if time.Now().After(subscription.EndDate) {
		render.Error(w, r, "auth", http.StatusUnauthorized, "Your subscription has expired, please renew your subscription.")
		return
	}

	// TODO: Add Plan and Subscription data to the generated token
	// so that we can directly access that using clains
	accessToken, err := generateAccessToken(user, subscription)
	if err != nil {
		render.Error(w, r, "auth", http.StatusInternalServerError, "Failed to generate access token.")
		return
	}

	render.JSON(w, http.StatusOK, struct {
		Token string `json:"access_token"`
		Code  int    `json:"code"`
	}{
		Token: accessToken,
		Code:  http.StatusOK,
	})
}

func generateAccessToken(user *generated.User, subscription *generated.Subscription) (string, error) {
	// TODO: Add permissions
	claims := models.CustomClaims{
		AccountID:      user.AccountID,
		PlanID:         subscription.Edges.Plan.ID,
		SubscriptionID: subscription.ID,
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
