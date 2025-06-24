package queries

import (
	"context"
	"log"

	library "github.com/anicse37/Library_Management/Backend"
)

func InsertBooks(ctx context.Context, db library.Database, book library.Book) {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description,year)
	VALUES (?,?,?,?);`, book.Name, book.Author, book.Description, book.Year); err != nil {
		log.Fatalf("Error While Inserting: %v", err)
	}
}
