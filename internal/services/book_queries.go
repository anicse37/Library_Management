package queries

import (
	"context"

	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func GetAllBooks(ctx context.Context, db models.Database) (models.ListBooks, error) {
	books, err := librarySQL.GetAllBooks(ctx, db)
	return books, err
}
func GetAllBorrowedBooks(ctx context.Context, db models.Database, userid string) (models.ListBooks, error) {
	books, err := librarySQL.GetAllBorrowedBooks(ctx, db, userid)
	return books, err
}

func AddBooks(ctx context.Context, db models.Database, book models.Book) error {
	err := librarySQL.InsertBooks(ctx, db, book)
	return err
}
func RemoveBooks(ctx context.Context, db models.Database, id int) error {
	err := librarySQL.DeleteBooks(ctx, db, id)
	return err
}

func AddBorrowBook(ctx context.Context, db models.Database, book models.Borrowed_Book) error {
	err := librarySQL.InsertBorrowedBooks(ctx, db, book)
	return err
}

func ReturnBorrowedBook(ctx context.Context, db models.Database, book string) error {
	_, err := db.DB.QueryContext(ctx, `UPDATE books SET available_no = available_no +1 WHERE id = ?`, book)
	if err != nil {
		return err
	}

	err = librarySQL.DeleteBorrowedBook(ctx, db, book)
	return err
}
func BorrowedBooks(ctx context.Context, db models.Database, user_id string) (models.ListBorrowedBookDisplay, error) {
	var books models.ListBorrowedBookDisplay
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ? ORDER BY returned_date IS NULL,borrow_date;", user_id)
	if err != nil {
		return books, err
	}
	book := models.BorrowedBookDisplay{}
	for result.Next() {
		result.Scan(&book.BorrowID, &book.UserId, &book.BookID, &book.BorrowDate, &book.ReturnedDate)
		result2 := db.DB.QueryRowContext(ctx, "SELECT name, author FROM books WHERE id = ? ;", book.BookID)
		result2.Scan(&book.BookName, &book.Author)
		books = append(books, book)
	}
	return books, nil
}
