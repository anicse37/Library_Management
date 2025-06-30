package books

import (
	"context"
	"net/http"
	"strconv"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	"github.com/anicse37/Library_Management/internal/search"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func BooksHandle(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			searchQuerry := r.URL.Query().Get("search")
			var books models.ListBooks
			var err error
			if searchQuerry != "" {
				books, err = search.SearchBook(ctx, db, searchQuerry)
			} else {
				books, err = queries.GetAllBooks(ctx, db)
			}
			if err != nil {
				errors_package.SetError(err)
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
				Query: searchQuerry,
				Role:  role,
			}
			template.RenderTemplate(w, "books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func AddBooksHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			template.RenderTemplate(w, "add_books.html", nil)
			return
		}
		r.ParseForm()
		name := r.FormValue("title")
		author := r.FormValue("author")
		year := r.FormValue("year")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		year1, _ := strconv.Atoi(year)
		quantity1, _ := strconv.Atoi(quantity)

		book := models.Book{
			Name:        name,
			Author:      author,
			Year:        year1,
			Description: description,
			Available:   quantity1,
		}

		queries.AddBooks(ctx, db, book)
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
func RemoveBooksHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		book, _ := strconv.Atoi(r.FormValue("book_id"))
		queries.RemoveBooks(ctx, db, book)
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
