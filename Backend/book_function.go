package library

import (
	"context"
	"database/sql"
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

func GetAllBorrowedBooks(ctx context.Context, db Database, userid string) ListBooks {
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ?;", userid)
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	var borrowed ListBooks
	borrowed_books := ScanBorrowedBooks(result)
	for _, j := range borrowed_books.Borrowed_Books {
		borrowed.Book = append(borrowed.Book, GetSingleBook(ctx, db, j.Book_id))
	}
	return borrowed
}

func GetSingleBook(ctx context.Context, db Database, book_id int) Book {
	result, _ := db.DB.QueryContext(ctx, "SELECT * FROM books where id = ?;", book_id)
	book := Book{}
	for result.Next() {
		result.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
	}
	return book

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

func ScanBorrowedBooks(rows *sql.Rows) ListBorrowed_Books {
	var all ListBorrowed_Books
	for rows.Next() {
		var b Borrowed_Book
		if err := rows.Scan(&b.Id, &b.User_id, &b.Book_id, &b.Borrow_Date, &b.Returned_Date); err != nil {
			continue
		}
		all.Borrowed_Books = append(all.Borrowed_Books, b)
	}
	return all
}
