package comment

import (
	"context"
	"database/sql"
	"log/slog"
	"task-tracker-service/internal/mappers/entitymap"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/helpers"
	"task-tracker-service/internal/types/entities"
	"task-tracker-service/internal/types/models"

	sq "github.com/Masterminds/squirrel"
)

type CommentDB interface {
	CreateComment(ctx context.Context, comment *entities.Comment) (*models.CommentModel, error)
}

type commentStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) CommentDB {
	return &commentStorage{
		db:     db,
		logger: logger,
	}
}

func (s *commentStorage) CreateComment(ctx context.Context, createComment *entities.Comment) (*models.CommentModel, error) {
	insertQuery, args, err := sq.Insert(dbconsts.TableComments).
		Columns(dbconsts.ColumnCommentTaskID, dbconsts.ColumnCommentAuthorID, dbconsts.ColumnCommentText).
		Values(createComment.TaskID, createComment.AuthorID, createComment.Text).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, insertQuery, args...)

	var comment entities.Comment

	err = row.Scan(&comment.ID, &comment.TaskID, &comment.AuthorID, &comment.Text, &comment.DateTime)
	if err != nil {
		return nil, helpers.CatchPQErrors(err)
	}

	return entitymap.MapToCommentModel(&comment), nil
}
