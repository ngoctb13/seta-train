package auth

import (
	"context"
	"net/http"
)

func AuthContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "Authorization", r.Header.Get("Authorization"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
