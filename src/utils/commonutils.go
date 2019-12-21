package utils

import (
	"os"
)

// GetEnv returns env value for key or fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
