package librarySQL

import (
	"context"
	"database/sql"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	"github.com/anicse37/Library_Management/internal/models"
)

func GetAllBooks(ctx context.Context, db models.Database) (models.ListBooks, error) {
	books := models.ListBooks{}
	result, err := db.DB.QueryContext(ctx, "SELECT *FROM books;")
	if err != nil {
		return books, errors_package.ErrorGettingBooks
	}
	books = ScanBooks(result)
	return books, nil
}

func GetAllBorrowedBooks(ctx context.Context, db models.Database, userid string) (models.ListBooks, error) {
	var borrowed models.ListBooks
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ?;", userid)
	if err != nil {
		return borrowed, errors_package.ErrorGettingBooks
	}
	borrowed_books := ScanBorrowedBooks(result)
	for _, j := range borrowed_books {
		book, err := GetSingleBook(ctx, db, j.Book_id)
		if err != nil {
			return borrowed, errors_package.ErrorGettingBooks
		}
		borrowed = append(borrowed, book)
	}
	return borrowed, nil
}

func GetSingleBook(ctx context.Context, db models.Database, book_id int) (models.Book, error) {
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM books where id = ?;", book_id)
	book := models.Book{}
	if err != nil {
		return book, errors_package.ErrorGettingBooks
	}
	for result.Next() {
		result.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
	}
	return book, nil

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
