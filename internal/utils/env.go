package utils

import "os"

// GetEnvAny Get Env From Any
func GetEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}

// GetEnvAnyWithDefault Get Env With Default
func GetEnvAnyWithDefault(defaultEnv string, names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return defaultEnv
}
