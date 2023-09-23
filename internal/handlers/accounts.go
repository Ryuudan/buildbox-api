package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/Pyakz/buildbox-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler struct {
	accountService services.AccountService
	userService    services.UserService
}

func NewAccountHandler(accountService services.AccountService, userService services.UserService) *AccountHandler {
	log.Println("✅ Accounts Handler Initialized")
	return &AccountHandler{
		accountService: accountService,
		userService:    userService,
	}
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	validate := utils.Validator()

	var account models.RegisterAccountStruct
	var validationErrors []utils.ValidationErrorDetails

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		utils.RenderError(w, r, "accounts", http.StatusBadRequest, err.Error())
		return
	}

	// Struct level validation
	if err := validate.Struct(account); err != nil {
		utils.Validate(w, r, err)
		return
	}

	user, _ := a.userService.GetUserByEmail(r.Context(), account.Email)

	if user != nil {
		validationErrors = append(validationErrors, utils.ValidationErrorDetails{
			Field:   "email",
			Message: "email already exists, please try another one",
		})
	}

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		utils.CustomValidationError(w, r, validationErrors)
		return
	}

	newAccount, err := a.accountService.CreateAccount(r.Context(), &generated.Account{
		Name:        account.Name,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
	})

	if err != nil {
		utils.RenderError(w, r, "accounts", http.StatusInternalServerError, err.Error())
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.RenderError(w, r, "users", http.StatusBadRequest, err.Error())
		return
	}

	// Creater user for the account
	newUser, err := a.userService.CreateUser(r.Context(), &generated.User{
		AccountID:  newAccount.ID,
		Email:      account.Email,
		Password:   string(password),
		FirstName:  account.FirstName,
		LastName:   account.LastName,
		MiddleName: account.MiddleName,
	})

	println(newUser)

	if err != nil {
		utils.RenderError(w, r, "accounts", http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	utils.RenderJSON(w, http.StatusCreated, newAccount)
}

func (a *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.accountService.GetAccounts(r.Context())
	if err != nil {
		utils.RenderError(w, r, "accounts", http.StatusInternalServerError, err.Error())
		return
	}

	utils.RenderJSON(w, http.StatusOK, accounts)
}
