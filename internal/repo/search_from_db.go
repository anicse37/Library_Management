package librarySQL

import (
	"context"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	"github.com/anicse37/Library_Management/internal/models"
)

func SearchWithRole(ctx context.Context, db models.Database, role string, keyword string) (models.ListUser, error) {
	users := models.ListUser{}
	user := models.User{}

	query := `SELECT * FROM user WHERE (name LIKE ? OR id LIKE ?) AND role = ?`
	likePattern := "%" + keyword + "%"
	res, err := db.DB.QueryContext(ctx, query, likePattern, likePattern, role)
	if err != nil {
		return users, errors_package.ErrorScanningUsers
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users = append(users, user)
	}
	return users, nil
}
func SearchBook(ctx context.Context, db models.Database, keyword string) (models.ListBooks, error) {
	books := models.ListBooks{}
	book := models.Book{}
	query := `SELECT * FROM books WHERE name LIKE ? OR author LIKE ?;`
	likePattern := "%" + keyword + "%"
	rows, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		return books, errors_package.ErrorGettingBooks
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books = append(books, book)
	}
	return books, nil
}
