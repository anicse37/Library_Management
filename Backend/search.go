package library

import (
	"context"
	"log"

	"github.com/anicse37/Library_Management/internal/models"
)

func SearchWithRole(ctx context.Context, db models.Database, role string, keyword string) (models.ListUser, error) {
	users := models.ListUser{}
	user := models.User{}

	query := `SELECT * FROM user WHERE (name LIKE ? OR id LIKE ?) AND role = ?`
	likePattern := "%" + keyword + "%"
	res, err := db.DB.QueryContext(ctx, query, likePattern, likePattern, role)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return users, err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users = append(users, user)
	}
	return users, nil
}

func SearchBook(ctx context.Context, db models.Database, keyword string) models.ListBooks {
	books := models.ListBooks{}
	book := models.Book{}

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
		books = append(books, book)
	}
	return books
}
func SearchBorrowedBook(ctx context.Context, db models.Database, keyword string) models.ListBooks {
	books := models.ListBooks{}
	book := models.Book{}

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
		books = append(books, book)
	}
	return books
}
