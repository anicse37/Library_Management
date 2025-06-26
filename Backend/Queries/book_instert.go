package queries

import (
	"context"
	"errors"
	"log"

	library "github.com/anicse37/Library_Management/Backend"
)

var (
	ErrorCanNotRemoveBook = errors.New("can not remove books")
	ErrorWhileInserting   = errors.New("errors while inserting")
)

func InsertBooks(ctx context.Context, db library.Database, book library.Book) {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description,year,available_no)
	VALUES (?,?,?,?,?);`, book.Name, book.Author, book.Description, book.Year, book.Available); err != nil {
		log.Fatalf("Error While Inserting: %v", err)
	}
}
func InsertBorrowedBooks(ctx context.Context, db library.Database, book library.Borrowed_Book) error {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO borrowed_books (user_id,book_id,borrow_date)
	VALUES (?,?,?);`, book.User_id, book.Book_id, book.Borrow_Date); err != nil {
		return ErrorWhileInserting
	}
	if _, err := db.DB.ExecContext(ctx, "UPDATE books SET available_no = available_no - 1 WHERE (id = ?) AND available_no >0", book.Book_id); err != nil {
		return ErrorCanNotRemoveBook
	}
	return nil
}
