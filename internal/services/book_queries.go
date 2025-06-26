package queries

import (
	"context"
	"fmt"

	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func GetAllBooks(ctx context.Context, db models.Database) models.ListBooks {
	books := librarySQL.GetAllBooks(ctx, db)
	return books
}
func GetAllBorrowedBooks(ctx context.Context, db models.Database, userid string) models.ListBooks {
	books := librarySQL.GetAllBorrowedBooks(ctx, db, userid)
	fmt.Println(books)

	return books
}
func RemoveBooks(ctx context.Context, db models.Database, id int) {
	db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE id = ?`, id)
	db.DB.ExecContext(ctx, `DELETE FROM books WHERE id = ?`, id)
}
func AddBooks(ctx context.Context, db models.Database, book models.Book) {
	librarySQL.InsertBooks(ctx, db, book)
}
func BorrowBook(ctx context.Context, db models.Database, book models.Borrowed_Book) {
	librarySQL.InsertBorrowedBooks(ctx, db, book)
}
func BorrowedBooks(ctx context.Context, db models.Database, user_id string) models.ListBorrowedBookDisplay {
	var books models.ListBorrowedBookDisplay
	result, _ := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ?;", user_id)
	book := models.BorrowedBookDisplay{}
	for result.Next() {
		result.Scan(&book.BookID, &book.UserId, &book.BookID, &book.BorrowDate, &book.ReturnedDate)
		result2 := db.DB.QueryRowContext(ctx, "SELECT name, author FROM books WHERE id = ? ;", book.BookID)
		result2.Scan(&book.BookName, &book.Author)
		books = append(books, book)
	}
	return books
}
