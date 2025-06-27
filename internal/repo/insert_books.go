package librarySQL

import (
	"context"

	"github.com/anicse37/Library_Management/internal/models"
)

func InsertBooks(ctx context.Context, db models.Database, book models.Book) error {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description,year,available_no)
	VALUES (?,?,?,?,?);`, book.Name, book.Author, book.Description, book.Year, book.Available); err != nil {
		return models.ErrorWhileInserting
	}
	return nil
}
func InsertBorrowedBooks(ctx context.Context, db models.Database, book models.Borrowed_Book) error {
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return models.ErrorWhileInserting
	}

	_, err = tx.ExecContext(ctx, `
    INSERT INTO borrowed_books (user_id, book_id, borrow_date)
    VALUES (?, ?, ?)`,
		book.User_id, book.Book_id, book.Borrow_Date,
	)
	if err != nil {
		tx.Rollback()
		return models.ErrorWhileInserting
	}

	_, err = tx.ExecContext(ctx, `
    UPDATE books
    SET available_no = available_no - 1
    WHERE id = ? AND available_no > 0`,
		book.Book_id,
	)
	if err != nil {
		tx.Rollback()
		return models.ErrorWhileInserting
	}
	if err = tx.Commit(); err != nil {
		return models.ErrorWhileInserting
	}
	return nil
}
