package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
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

type ValidationErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type validationError struct {
	Code    int                      `json:"code"`
	Type    string                   `json:"type"`
	Details []ValidationErrorDetails `json:"details"`
}

type validationErrorResponse struct {
	Error validationError `json:"error"`
}

func validationErrorMessage(err validator.FieldError) string {
	field := err.Field()

	switch tag := err.Tag(); tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s should be at least %s characters", field, err.Param())
	case "max":
		return fmt.Sprintf("%s should be at most %s characters", field, err.Param())
	case "lte":
		fmt.Println(err.Value())
		return fmt.Sprintf("%s should be less than or equal %s", field, err.Param())
	case "gte":
		fmt.Println(err.Value())
		return fmt.Sprintf("%s should be greater than or equal %s", field, err.Param())
	case "email":
		return fmt.Sprintf("%s is not a valid email address", field)
	case "url":
		return fmt.Sprintf("%s is not a valid URL", field)
	case "oneof":
		return fmt.Sprintf("%s should be one of [%s]", field, err.Param())
	case "alpha":
		return fmt.Sprintf("%s should only contain alphabetic characters", field)
	case "alphanum":
		return fmt.Sprintf("%s should only contain alphanumeric characters", field)
	case "numeric":
		return fmt.Sprintf("%s should be a numeric value", field)
	case "number":
		return fmt.Sprintf("%s should be a number value", field)
	default:
		return fmt.Sprintf("Validation failed for %s field", field)
	}
}

// ideal for form validation, or form errorrs
// like setError in react-hooks-form
func Validate(w http.ResponseWriter, err error) {
	var details []ValidationErrorDetails
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errorType := validationErrorMessage(err)
		details = append(details, ValidationErrorDetails{Field: field, Message: errorType})
	}

	errorResponse := validationErrorResponse{
		Error: validationError{
			Code:    http.StatusBadRequest,
			Type:    "validation_error",
			Details: details,
		},
	}

	response, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// ideal for form validation, or form errorrs
// like setError in react-hooks-form
func CustomValidationError(w http.ResponseWriter, details []ValidationErrorDetails) {

	errorResponse := validationErrorResponse{
		Error: validationError{
			Code:    http.StatusBadRequest,
			Type:    "validation_error",
			Details: details,
		},
	}

	response, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
