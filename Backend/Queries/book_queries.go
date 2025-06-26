package queries

import (
	"context"
	"fmt"

	library "github.com/anicse37/Library_Management/Backend"
)

func GetAllBooks(ctx context.Context, db library.Database) library.ListBooks {
	books := library.GetAllBooks(ctx, db)
	return books
}
func GetAllBorrowedBooks(ctx context.Context, db library.Database, userid string) library.ListBooks {
	books := library.GetAllBorrowedBooks(ctx, db, userid)
	fmt.Println(books)
	return books
}
func RemoveBooks(ctx context.Context, db library.Database, id int) {
	db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE id = ?`, id)
	db.DB.ExecContext(ctx, `DELETE FROM books WHERE id = ?`, id)
}
func AddBooks(ctx context.Context, db library.Database, book library.Book) {
	InsertBooks(ctx, db, book)
}
func BorrowBook(ctx context.Context, db library.Database, book library.Borrowed_Book) {
	InsertBorrowedBooks(ctx, db, book)
}
func BorrowedBooks(ctx context.Context, db library.Database, user_id string) library.ListBorrowedBookDisplay {
	var books library.ListBorrowedBookDisplay
	result, _ := db.DB.QueryContext(ctx, "SELECT * FROM borrowed_books WHERE user_id = ?;", user_id)
	book := library.BorrowedBookDisplay{}
	for result.Next() {
		result.Scan(&book.BookID, &book.UserId, &book.BookID, &book.BorrowDate, &book.ReturnedDate)
		result2 := db.DB.QueryRowContext(ctx, "SELECT name, author FROM books WHERE id = ? ;", book.BookID)
		result2.Scan(&book.BookName, &book.Author)
		books = append(books, book)
	}
	return books
}
