package librarySQL

import (
	"context"

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
