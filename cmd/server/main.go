package main

import (
	"context"
	"fmt"

	"github.com/TutotialEdge/go-rest-api/internal/comment"
	"github.com/TutotialEdge/go-rest-api/internal/db"
)

// Run - initiation and startup
func Run() error {
	fmt.Println("starting the application")

	db_, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}
	if err := db_.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	commentService := comment.NewService(db_)

	commentService.PostComment(context.Background(),
		comment.Comment{
			ID:     "ea43fb26-01c4-4c4d-9ec4-4b9061bb21f0",
			Slug:   "manual-test",
			Author: "Anand",
			Body:   "Hello World",
		},
	)

	fmt.Println(commentService.GetComment(
		context.Background(),
		"ea43fb26-01c4-4c4d-9ec4-4b9061bb21f0",
	))

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
