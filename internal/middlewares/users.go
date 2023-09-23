package middlewares

import (
	"net/http"

	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/internal/services"
	"github.com/golang-jwt/jwt"
)

type UserMiddleware struct {
	userService services.UserService
}

func NewUserMiddleware(userService services.UserService) *UserMiddleware {
	return &UserMiddleware{
		userService: userService,
	}
}

func (u *UserMiddleware) ValidUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(models.ContextKeyClaims).(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := u.userService.GetUserById(r.Context(), int(claims["user_id"].(float64)))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
