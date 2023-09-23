package accounts

import (
	"encoding/json"
	"net/http"

	models "github.com/Pyakz/buildbox-api/ent/generated"
	"github.com/Pyakz/buildbox-api/utils"
)

type AccountHandler struct {
	accountService AccountService
}

func NewAccountHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	validate := utils.Validator()

	var account models.Account
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

	// If there are validation errors, return a custom validation error response
	if len(validationErrors) > 0 {
		utils.CustomValidationError(w, r, validationErrors)
		return
	}

	newAccount, err := a.accountService.CreateAccount(r.Context(), &account)

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
