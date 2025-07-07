package books

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func BorrowedBooksHandle(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			session, _ := session.Store.Get(r, "very-secret-key")
			userid, _ := session.Values[models.SessionKeyUserId].(string)

			books, err := queries.GetAllBorrowedBooks(ctx, db, userid)
			if err != nil {
				errors_package.SetError(err)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			}

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

func BorrowHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		session, _ := session.Store.Get(r, "very-secret-key")
		userID, _ := session.Values[models.SessionKeyUserId].(string)

		book_id, _ := strconv.Atoi(r.FormValue("book_id"))
		book := models.Borrowed_Book{
			User_id:     userID,
			Book_id:     book_id,
			Borrow_Date: time.Now(),
		}
		err := queries.AddBorrowBook(ctx, db, book)
		if err != nil {
			errors_package.SetError(err)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/your_books", http.StatusSeeOther)
		BorrowedBooksHandler(ctx, db)
	}
}
func BorrowedBooksHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		session, _ := session.Store.Get(r, "very-secret-key")
		user_id, _ := session.Values[models.SessionKeyUserId].(string)
		role, _ := session.Values[models.SessionKeyRole].(string)

		books, err := queries.BorrowedBooks(ctx, db, user_id)
		if err != nil {
			http.Redirect(w, r, "/books?msg=error_in_borrowed_books", http.StatusSeeOther)
		}

		data := struct {
			Book models.ListBorrowedBookDisplay
			Role string
		}{
			Book: books,
			Role: role,
		}
		template.RenderTemplate(w, "borrowed_books.html", data)
	}
}

func ReturnBookHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		book := r.FormValue("book_id")
		fmt.Println(book)
		err := queries.ReturnBorrowedBook(ctx, db, book)
		if err != nil {
			http.Redirect(w, r, "/your_books?msg=return_error", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/your_books", http.StatusSeeOther)
	}
}
