package books

import (
	"context"
	"net/http"
	"strconv"
	"time"

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
				//do something
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
		queries.AddBorrowBook(ctx, db, book)
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

		books := queries.BorrowedBooks(ctx, db, user_id)
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
