package utils

import (
	"fmt"
	"os"
)

func LoadEnvironmentVariables() error {

	requiredVariables := []string{"PORT", "POSTGRES_CONNECTION_STRING", "JWT_SECRET", "REDIS_PORT"}
	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			return fmt.Errorf("missing required environment variable: %s", variable)
		}
	}

	return nil
}
