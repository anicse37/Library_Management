package books

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func BooksHandle(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var books library.ListBookJSON
			if search != "" {
				books = db.SearchBook(ctx, search)
			} else {
				books = db.GetAllBooks(ctx)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[library.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Book  library.ListBookJSON
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
