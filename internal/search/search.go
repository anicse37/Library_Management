package search

import (
	"context"

	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func SearchUsers(ctx context.Context, db models.Database, role string, keyword string) (models.ListUser, error) {
	users, err := librarySQL.SearchWithRole(ctx, db, role, keyword)
	return users, err
}

func SearchBook(ctx context.Context, db models.Database, keyword string) (models.ListBooks, error) {
	book, err := librarySQL.SearchBook(ctx, db, keyword)
	return book, err
}
