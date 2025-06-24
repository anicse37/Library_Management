package library

import (
	"context"
	"log"
)

func (db *Database) SearchUsers(ctx context.Context, keyword string) ListUser {
	users := ListUser{}
	user := User{}
	query := `SELECT * FROM user WHERE name LIKE ? OR id LIKE ?;`
	likePattern := "%" + keyword + "%"
	res, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return users
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users
}

func (db *Database) SearchBook(ctx context.Context, keyword string) ListBooks {
	books := ListBooks{}
	book := Book{}

	query := `SELECT * FROM books WHERE name LIKE ? OR author LIKE ?;`
	likePattern := "%" + keyword + "%"
	rows, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return books
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books.Book = append(books.Book, book)
	}
	return books
}
func (db *Database) SearchBorrowedBook(ctx context.Context, keyword string) ListBooks {
	books := ListBooks{}
	book := Book{}

	query := `SELECT * FROM borrowed_books WHERE name LIKE ? OR author LIKE ?;`
	likePattern := "%" + keyword + "%"
	rows, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return books
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books.Book = append(books.Book, book)
	}
	return books
}
