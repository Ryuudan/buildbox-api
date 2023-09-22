package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

func GenerateID() string {
	return uuid.New().String()
}

// RenderJSON writes response from struct to JSON.
func RenderJSON(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// RenderError writes an error response as JSON to the client.
func RenderError(w http.ResponseWriter, errType string, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Create a buffer to hold the JSON encoding.
	var buf []byte
	data := map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"type":    errType,
			"message": message,
		},
	}

	// Encode the data into the buffer.
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the buffer's content to the response writer.
	if _, err := w.Write(buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
