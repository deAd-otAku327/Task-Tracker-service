package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (hub *middlewareHub) ResponseTimeLimit() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), hub.params.RespTimeLimit)
			defer cancel()

			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}
