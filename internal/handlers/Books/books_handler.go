package books

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func BooksHandle(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var books models.ListBooks
			if search != "" {
				books = library.SearchBook(ctx, db, search)
			} else {
				books = queries.GetAllBooks(ctx, db)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[models.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Book  models.ListBooks
				Query string
				Role  string
			}{
				Book:  books,
				Query: search,
				Role:  role,
			}
			template.RenderTemplate(w, "books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func BorrowedBooksHandle(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var books models.ListBooks

			session, _ := session.Store.Get(r, "very-secret-key")
			userid, _ := session.Values[models.SessionKeyUserId].(string)

			books = queries.GetAllBorrowedBooks(ctx, db, userid)

			role := "user"
			if rRole, ok := session.Values[models.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Book models.ListBooks
				Role string
			}{
				Book: books,
				Role: role,
			}
			template.RenderTemplate(w, "borrowed_books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
