package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (commentRepositoryImpl *commentRepositoryImpl) Create(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comment (username, message) VALUES (?, ?);"

	result, err := commentRepositoryImpl.DB.ExecContext(ctx, query, comment.Username, comment.Message)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int(id)

	return comment, nil
}

func (commentRepositoryImpl *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, username, message FROM comment;"

	var comments []entity.Comment

	rows, err := commentRepositoryImpl.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		comment := entity.Comment{}

		err := rows.Scan(&comment.Id, &comment.Username, &comment.Message)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (commentRepositoryImpl *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comment, error) {
	query := "SELECT id, username, message FROM comment WHERE id = ? LIMIT 1;"

	comment := entity.Comment{}

	rows, err := commentRepositoryImpl.DB.QueryContext(ctx, query, id)

	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Username, &comment.Message)

		if err != nil {
			return comment, err
		}

		return comment, nil
	} else {
		return comment, errors.New("Comment " + strconv.Itoa(id) + " Not Found")
	}
}
