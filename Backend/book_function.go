package library

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func GetAllBooks(ctx context.Context, db Database) ListBooks {

	result, err := db.DB.QueryContext(ctx, "SELECT *FROM books;")
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	books := ScanBooks(result)
	return books
}
func GetBooksById(ctx context.Context, db Database, id int) {

}

func GetAllBorrowedBooks(ctx context.Context, db Database) ListBooks {
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books;")
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	borrowed_books := ScanBooks(result)
	for i, j := range borrowed_books.Book {
		fmt.Printf("no: %v, books:%v", i, j)
	}
	return borrowed_books
}

// Below are helper functions
func ScanBooks(result *sql.Rows) ListBooks {
	books := ListBooks{}
	book := Book{}
	for result.Next() {
		result.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books.Book = append(books.Book, book)
	}
	return books
}
