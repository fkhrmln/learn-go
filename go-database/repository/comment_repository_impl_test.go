package repository

import (
	"context"
	"fmt"
	"go-database/database"
	"go-database/entity"
	"testing"
)

func TestCreateComment(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Username: "Fakhri Maulana Ihsan",
		Message:  "Hello World",
	}

	result, err := commentRepository.CreateComment(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
