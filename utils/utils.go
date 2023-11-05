package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

func GenerateID() string {
	return uuid.New().String()
}

func VersionMiddleware(version string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-API-Version", version)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func ConsoleLog(message interface{}) {
	pp.Println(message)
}

// StringToInt returns the integer value of a string or an error if the conversion fails.
func StringToInt(str string) (int, error) {

	if str == "" {
		return 0, errors.New("provided string is empty")
	}

	intValue, err := strconv.Atoi(str)

	if err != nil {
		return 0, err
	}

	return intValue, nil
}

var DEFAULT_OWNER_PERMISSIONS = []string{
	"account:read",
	"account:update",
	"account:delete",
	"account:settings",
	"account:billing",
	"account:subscription",
	"users:read",
	"users:create",
	"users:update",
	"users:update_others",
	"users:delete",
	"issues:read",
	"issues:create",
	"issues:update",
	"issues:delete",
	"plans:read",
	"plans:create",
	"plans:update",
	"plans:delete",
	"projects:read",
	"projects:create",
	"projects:update",
	"projects:delete",
	"milestones:read",
	"milestones:create",
	"milestones:update",
	"milestones:delete",
	"roles:read",
	"roles:create",
	"roles:update",
	"roles:delete",
	"service_providers:read",
	"service_providers:create",
	"service_providers:update",
	"service_providers:delete",
	"tasks:read",
	"tasks:create",
	"tasks:update",
	"tasks:delete",
	"teams:read",
	"teams:create",
	"teams:update",
	"teams:delete",
}
