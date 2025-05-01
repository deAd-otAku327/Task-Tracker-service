package middleware

import (
	"log/slog"
	"task-tracker-service/internal/tokenizer"
	"time"

	"github.com/gorilla/mux"
)

type contextKey int8

const (
	CookieName = "token"

	RequestIDKey contextKey = iota
	UserIDKey
)

type Middleware interface {
	Auth() mux.MiddlewareFunc
	Logging() mux.MiddlewareFunc
	RpsLimit() mux.MiddlewareFunc
	ResponseTimeLimit() mux.MiddlewareFunc
}

type MwParams struct {
	RpsLimit      int
	RespTimeLimit time.Duration
}

type middlewareHub struct {
	tokenizer tokenizer.Tokenizer
	logger    *slog.Logger
	params    MwParams
}

func New(tok tokenizer.Tokenizer, log *slog.Logger, params MwParams) Middleware {
	return &middlewareHub{
		tokenizer: tok,
		logger:    log,
		params:    params,
	}
}
