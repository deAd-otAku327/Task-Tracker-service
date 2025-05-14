package middleware

import (
	"context"
	"net/http"
	"task-tracker-service/internal/controller/_shared/apierrors"
	"task-tracker-service/internal/controller/_shared/responser"
	"task-tracker-service/internal/mappers/errmap"

	"github.com/gorilla/mux"
)

func (hub *middlewareHub) Auth() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie(CookieName)
			if err != nil {
				responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrAuthenticationFailed, http.StatusForbidden))
				return
			}

			token, err := hub.tokenizer.VerifyToken(tokenCookie.Value)
			if err != nil {
				responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrAuthenticationFailed, http.StatusForbidden))
				return
			}

			userID, err := token.Claims.GetSubject()
			if err != nil {
				responser.MakeErrorResponseJSON(w, errmap.MapToErrorResponse(apierrors.ErrAuthenticationFailed, http.StatusForbidden))
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserIDKey, userID)))
		})
	}
}
