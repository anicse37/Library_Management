package books

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func BooksHandle(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var books library.ListBooks
			if search != "" {
				books = library.SearchBook(ctx, db, search)
			} else {
				books = queries.GetAllBooks(ctx, db)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[library.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Book  library.ListBooks
				Query string
				Role  string
			}{
				Book:  books,
				Query: search,
				Role:  role,
			}
			server.RenderTemplate(w, "books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func BorrowedBooksHandle(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var books library.ListBooks

			session, _ := session.Store.Get(r, "very-secret-key")
			userid, _ := session.Values[library.SessionKeyUserId].(string)

			books = queries.GetAllBorrowedBooks(ctx, db, userid)

			role := "user"
			if rRole, ok := session.Values[library.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Book library.ListBooks
				Role string
			}{
				Book: books,
				Role: role,
			}
			server.RenderTemplate(w, "borrowed_books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
