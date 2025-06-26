package librarySQL

import (
	"context"
	"database/sql"
	"log"

	"github.com/anicse37/Library_Management/internal/models"
)

func GetAllBooks(ctx context.Context, db models.Database) models.ListBooks {
	result, err := db.DB.QueryContext(ctx, "SELECT *FROM books;")
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	books := ScanBooks(result)
	return books
}

func GetAllBorrowedBooks(ctx context.Context, db models.Database, userid string) models.ListBooks {
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ?;", userid)
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	var borrowed models.ListBooks
	borrowed_books := ScanBorrowedBooks(result)
	for _, j := range borrowed_books {
		borrowed = append(borrowed, GetSingleBook(ctx, db, j.Book_id))
	}
	return borrowed
}

func GetSingleBook(ctx context.Context, db models.Database, book_id int) models.Book {
	result, _ := db.DB.QueryContext(ctx, "SELECT * FROM books where id = ?;", book_id)
	book := models.Book{}
	for result.Next() {
		result.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
	}
	return book

}

// Below are helper functions
func ScanBooks(result *sql.Rows) models.ListBooks {
	books := models.ListBooks{}
	book := models.Book{}
	for result.Next() {
		result.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books = append(books, book)
	}
	return books
}

func ScanBorrowedBooks(rows *sql.Rows) models.ListBorrowed_Books {
	var all models.ListBorrowed_Books
	for rows.Next() {
		var b models.Borrowed_Book
		if err := rows.Scan(&b.Id, &b.User_id, &b.Book_id, &b.Borrow_Date, &b.Returned_Date); err != nil {
			continue
		}
		all = append(all, b)
	}
	return all
}
