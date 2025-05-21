package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (hub *middlewareHub) Logging() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client := r.RemoteAddr

			rid := uuid.New().String()
			w.Header().Set("X-Request-ID", rid)

			hub.logger.Info(
				"Request:",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.String("client", client),
				slog.String("request_id", rid),
			)

			startReq := time.Now()

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), RequestIDKey, rid)))

			responseTime := time.Since(startReq).Milliseconds()

			hub.logger.Info(
				"Response:",
				slog.String("client", client),
				slog.String("resp_time", strconv.Itoa(int(responseTime))),
				slog.String("request_id", rid),
			)
		})
	}
}
