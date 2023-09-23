package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	log.Println("✅ User Handler Initialized")
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) GetUsersByAccount(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetUsersByAccountID(r.Context(), 1)
	if err != nil {
		utils.RenderError(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}
	utils.RenderJSON(w, http.StatusOK, users)
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	user, err := u.userService.GetUserById(r.Context(), 1)
	if err != nil {
		utils.RenderError(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}
	utils.RenderJSON(w, http.StatusOK, user)
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	validate := utils.Validator()

	var user models.CreateUserStruct
	var validationErrors []utils.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RenderError(w, r, "users", http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(user); err != nil {
		utils.Validate(w, r, err)
		return
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		utils.CustomValidationError(w, r, validationErrors)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RenderError(w, r, "users", http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := u.userService.CreateUser(r.Context(), &generated.User{
		AccountID: 1,
		Email:     user.Email,
		Password:  string(password),
	})

	if err != nil {
		utils.RenderError(w, r, "users", http.StatusInternalServerError, err.Error())
		return
	}

	utils.RenderJSON(w, http.StatusOK, newUser)
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	validate := utils.Validator()

	var credentials models.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.RenderError(w, r, "users", http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(credentials); err != nil {
		utils.Validate(w, r, err)
		return
	}

	user, err := u.userService.GetUserByEmail(r.Context(), credentials.Email)

	if err != nil {
		utils.CustomValidationError(w, r, []utils.ValidationErrorDetails{
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
		utils.CustomValidationError(w, r, []utils.ValidationErrorDetails{
			{
				Field:   "password",
				Message: "Invalid password, please try again.",
			},
		})
		return
	}

	accessToken, err := generateAccessToken(user)
	if err != nil {
		utils.RenderError(w, r, "auth", http.StatusInternalServerError, "Failed to generate access token.")
		return
	}

	utils.RenderJSON(w, http.StatusOK, struct {
		Token string `json:"access_token"`
		Code  int    `json:"code"`
	}{
		Token: accessToken,
		Code:  http.StatusOK,
	})
}

func generateAccessToken(user *generated.User) (string, error) {

	claims := models.CustomClaims{
		AccountID: user.AccountID,
		UserID:    user.ID,
		UserUUID:  user.UUID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
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
