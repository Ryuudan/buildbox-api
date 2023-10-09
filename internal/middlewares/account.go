package middlewares

import (
	"net/http"

	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

// This ensure that users from another account cannot access other accounts
func EnforceAccountAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user claims from the JWT token in the request context
		claims, ok := r.Context().Value(models.ContextKeyClaims).(jwt.MapClaims)
		if !ok {
			render.Error(w, r, http.StatusBadRequest, "failed to get user claims from context")
			return
		}

		id, err := utils.StringToInt(chi.URLParam(r, "id"))

		if err != nil {
			render.Error(w, r, http.StatusBadRequest, "Invalid ID")
			return
		}

		accountID := int(claims["account_id"].(float64))

		if id != accountID {
			render.Error(w, r, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
