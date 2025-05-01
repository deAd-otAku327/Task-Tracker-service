package user

import (
	"log/slog"
	"task-tracker-service/internal/storage/db"
	"task-tracker-service/internal/tokenizer"
	"task-tracker-service/pkg/cryptor"
)

type UserService interface {
}

type userService struct {
	storage   db.DB
	logger    *slog.Logger
	cryptor   cryptor.Cryptor
	tokenizer tokenizer.Tokenizer
}

func New(s db.DB, logger *slog.Logger, cryptor cryptor.Cryptor, tok tokenizer.Tokenizer) UserService {
	return &userService{
		storage:   s,
		logger:    logger,
		cryptor:   cryptor,
		tokenizer: tok,
	}
}
