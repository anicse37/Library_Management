package queries

import (
	"context"
	"fmt"

	library "github.com/anicse37/Library_Management/Backend"
)

func GetAllBooks(ctx context.Context, db library.Database) library.ListBooks {
	books := library.GetAllBooks(ctx, db)
	return books
}
func GetAllBorrowedBooks(ctx context.Context, db library.Database, userid string) library.ListBooks {
	books := library.GetAllBorrowedBooks(ctx, db, userid)
	fmt.Println(books)
	return books
}
