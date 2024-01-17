package repository

import (
	belajargolangdatabase "belajar_golang_database"
	"belajar_golang_database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 25)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
