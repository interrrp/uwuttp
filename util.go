// Provides common utilities.

package main

import "os"

// envOr returns the value of an environment variable, or a fallback if it
// doesn't exist.
func envOr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
