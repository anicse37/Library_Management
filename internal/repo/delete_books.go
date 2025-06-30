package librarySQL

import (
	"context"
	"time"

	"github.com/anicse37/Library_Management/internal/models"
)

func DeleteBooks(ctx context.Context, db models.Database, id int) error {
	if _, err := db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE id = ?`, id); err != nil {
		return models.ErrorWhileRemoveing
	}
	if _, err := db.DB.ExecContext(ctx, `DELETE FROM books WHERE id = ?`, id); err != nil {
		return models.ErrorWhileRemoveing
	}
	return nil
}

func DeleteBorrowedBook(ctx context.Context, db models.Database, id string) error {
	if _, err := db.DB.ExecContext(ctx, `UPDATE borrowed_books SET returned_date = ? WHERE id = ?`, id, time.Now()); err != nil {
		return models.ErrorWhileRemoveing
	}
	return nil
}
