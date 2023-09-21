package utils

import (
	"fmt"
	"os"
)

func LoadEnvironmentVariables() error {

	requiredVariables := []string{"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASS", "PORT"}
	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			return fmt.Errorf("missing required environment variable: %s", variable)
		}
	}

	return nil
}
