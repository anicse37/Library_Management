package librarySQL

import (
	"context"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	"github.com/anicse37/Library_Management/internal/models"
)

func DeleteBooks(ctx context.Context, db models.Database, id int) error {
	if _, err := db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE book_id = ?`, id); err != nil {
		return errors_package.ErrorWhileRemoveing
	}
	if _, err := db.DB.ExecContext(ctx, `DELETE FROM books WHERE id = ?`, id); err != nil {
		return errors_package.ErrorWhileRemoveing
	}
	return nil
}

func DeleteBorrowedBook(ctx context.Context, db models.Database, id string) error {
	if _, err := db.DB.ExecContext(ctx, `UPDATE borrowed_books SET returned_date = NOW() WHERE book_id = ?;`, id); err != nil {
		return errors_package.ErrorWhileRemoveing
	}
	return nil
}
