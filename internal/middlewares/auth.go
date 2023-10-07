package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Pyakz/buildbox-api/internal/models"
	"github.com/Pyakz/buildbox-api/utils/render"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			// Authorization header not found, return unauthorized status
			render.Error(w, r, http.StatusUnauthorized, "Please provide a valid token, or login to get one.")
			return
		}

		// Split the Authorization header to get the token value
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			// Invalid Authorization header format, return unauthorized status
			render.Error(w, r, http.StatusUnauthorized, "Please provide a valid token, or login to get one.")
			return
		}

		// Extract the token value from the Authorization header
		tokenString := headerParts[1]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the same secret key used to sign the token during generation
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			// Failed to parse JWT token, return unauthorized status
			render.Error(w, r, http.StatusUnauthorized, "Please provide a valid token, or login to get one.")
			return
		}

		// Check if the token is valid
		if !token.Valid {
			// Invalid token, return unauthorized status
			render.Error(w, r, http.StatusUnauthorized, "Please provide a valid token, or login to get one.")
			return
		}

		// Extract the custom claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			// Failed to extract claims, return unauthorized status
			render.Error(w, r, http.StatusUnauthorized, "Please provide a valid token, or login to get one.")
			return
		}

		// Add the custom claims to the request context
		ctx := context.WithValue(r.Context(), models.ContextKeyClaims, claims)

		// Call the next middleware or handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
