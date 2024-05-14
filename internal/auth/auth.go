package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API key
// from the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authentication info found")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}
	return parts[1], nil
}
