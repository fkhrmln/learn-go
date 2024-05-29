package repository

import (
	"context"
	"go-database/entity"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
