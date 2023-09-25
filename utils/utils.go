package utils

import (
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
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
