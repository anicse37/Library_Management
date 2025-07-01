package librarySQL

import (
	"context"
	"database/sql"
	"errors"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	"github.com/anicse37/Library_Management/internal/models"
)

func InsertBooks(ctx context.Context, db models.Database, book models.Book) error {
	var existingAvailable int
	query := `SELECT available_no FROM books 
	          WHERE name = ? AND author = ? AND year = ?`

	err := db.DB.QueryRowContext(ctx, query, book.Name, book.Author, book.Year).Scan(&existingAvailable)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Book does not exist → insert new record
			_, insertErr := db.DB.ExecContext(ctx,
				`INSERT INTO books (name, author, description, year, available_no)
				VALUES (?, ?, ?, ?, ?)`,
				book.Name, book.Author, book.Description, book.Year, book.Available)
			if insertErr != nil {
				return errors_package.ErrorWhileInserting
			}
			return nil
		}
		// Some other error occurred
		return err
	}

	// Book exists → update available_no
	newAvailable := existingAvailable + book.Available

	_, updateErr := db.DB.ExecContext(ctx,
		`UPDATE books 
		 SET available_no = ? 
		 WHERE name = ? AND author = ? AND year = ?`,
		newAvailable, book.Name, book.Author, book.Year)
	if updateErr != nil {
		return errors_package.ErrorWhileInserting
	}

	return nil
}
func InsertBorrowedBooks(ctx context.Context, db models.Database, book models.Borrowed_Book) error {
	var count int
	err := db.DB.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM borrowed_books
		WHERE user_id = ? AND book_id = ? AND returned_date IS NULL
	`, book.User_id, book.Book_id).Scan(&count)
	if err != nil {
		return errors_package.ErrorWhileInserting
	}

	if count > 0 {
		return errors_package.ErrorAlreadyBorrowed
	}

	var availableNo int
	err = db.DB.QueryRowContext(ctx, `
		SELECT available_no
		FROM books
		WHERE id = ?
	`, book.Book_id).Scan(&availableNo)
	if err != nil {
		return errors_package.ErrorWhileInserting
	}

	if availableNo <= 0 {
		return errors_package.ErrorNoBooksAvailable
	}
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return errors_package.ErrorWhileInserting
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO borrowed_books (user_id, book_id, borrow_date)
		VALUES (?, ?, ?)
	`,
		book.User_id,
		book.Book_id,
		book.Borrow_Date,
	)
	if err != nil {
		tx.Rollback()
		return errors_package.ErrorWhileInserting
	}

	result, err := tx.ExecContext(ctx, `
		UPDATE books
		SET available_no = available_no - 1
		WHERE id = ? AND available_no > 0
	`,
		book.Book_id,
	)
	if err != nil {
		tx.Rollback()
		return errors_package.ErrorWhileInserting
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return errors_package.ErrorWhileInserting
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return errors_package.ErrorNoBooksAvailable
	}

	if err = tx.Commit(); err != nil {
		return errors_package.ErrorWhileInserting
	}

	return nil
}
