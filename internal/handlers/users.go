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

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
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

	users, total, err := u.userService.GetUsers(r.Context(), queryParams, filters)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	render.JSON(w, http.StatusOK, &render.PaginatedResults{
		Results: users,
		Meta: render.GenerateMeta(
			total,
			queryParams,
			len(users),
		),
	})
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		render.Error(w, r, http.StatusBadRequest, constants.INVALID_FORMAT_ID)
		return
	}

	user, err := u.userService.GetUserByID(r.Context(), id)

	if err != nil {
		if generated.IsNotFound(err) {
			render.Error(w, r, http.StatusNotFound, "user not found")
			return
		} else {
			render.Error(w, r, http.StatusInternalServerError, err.Error())
			return
		}
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
		render.Error(w, r, http.StatusUnprocessableEntity, "Invalid JSON: "+err.Error())
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
		render.Error(w, r, http.StatusBadRequest, "failed to get user claims from context")
		return
	}

	// Assign the account ID from claims to the user
	user.AccountID = int(claims["account_id"].(float64))

	// Generate a salted and hashed password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, "failed to generate hashed password")
		return
	}

	// Replace the plaintext password with the hashed password
	user.Password = string(password)

	// Register the user in the system
	newUser, err := u.userService.RegisterUser(r.Context(), &user)

	if err != nil {
		render.Error(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the newly created user in JSON format
	render.JSON(w, http.StatusOK, newUser)
}
