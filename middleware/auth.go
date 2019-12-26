package middleware

import (
	"context"
	"github.com/nengkhoiba/go-framework-dda/utils"
	"log"
	"net/http"
)

type AuthKeys string

func (m *Middleware) RequireAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var userID int64
		var err error
		authHeader := r.Header.Get("Authorization")
		authToken, err := utils.HeaderToTokenString(authHeader)
		if err != nil {
			log.Println(err)
			utils.RespondWithError(w, http.StatusUnauthorized, err) // Directly returning err to user is harmless here. Custom token with harmless message
			return
		}

		// perform auth action

		ctx = context.WithValue(ctx, AuthKeys("user_id"), userID)
		ctx = context.WithValue(ctx, AuthKeys("auth_token"), authToken) //set value to context

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
