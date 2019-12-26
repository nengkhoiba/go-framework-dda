package utils

import (
	"github.com/nengkhoiba/go-framework-dda/errs"
	"strings"
)

// HeaderToTokenString returns token from the Authorization header
func HeaderToTokenString(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errs.NoHeader
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errs.NoBearer
	}
	authToken := authHeader[len("Bearer "):]
	return authToken, nil
}
