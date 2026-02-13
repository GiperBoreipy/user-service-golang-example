package presentation

import (
	"context"
	"net/http"

	"user_service/internal/application/data_objects"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		ctx := context.WithValue(r.Context(), data_objects.UserAuthTokenContextKey, data_objects.UserAuthToken(authHeader))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
