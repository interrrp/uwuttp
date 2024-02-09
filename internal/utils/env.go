package utils

import "os"

// EnvOr returns the value of an environment variable, or a fallback if it
// doesn't exist.
func EnvOr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
