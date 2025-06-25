package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func AllUsersHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var users library.ListUser
			if search != "" {
				users, _ = queries.SearchUsers(ctx, db, search)
			} else {
				users, _ = queries.GetAllUsers(ctx, db)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[library.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Users library.ListUser
				Query string
				Role  string
			}{
				Users: users,
				Query: search,
				Role:  role,
			}
			server.RenderTemplate(w, "all_users.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func AllAdminsHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			user1, _ := queries.GetAdminsNotApproved(ctx, db)
			user2, _ := queries.GetAdminsApproved(ctx, db)

			data := struct {
				Users1 library.ListUser
				Users2 library.ListUser
				Querry string
			}{
				Users1: user1,
				Users2: user2,
				Querry: search,
			}
			server.RenderTemplate(w, "manage_admins.html", data)
		default:
		}
	}
}

func ApproveHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		admin := r.FormValue("admin_id")
		fmt.Println(admin)
		queries.ApproveAdmin(ctx, db, admin)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/manage_admins", http.StatusSeeOther)
	}
}
func RemoveAdminHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		admin := r.FormValue("admin_id")
		queries.RemoveAdmin(ctx, db, admin)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/manage_admins", http.StatusSeeOther)
	}
}
func RemoveUserHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.FormValue("user_id")
		queries.RemoveUser(ctx, db, user)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/all_users", http.StatusSeeOther)
	}
}
func RemoveBooksHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		book := r.FormValue("book_id")
		queries.RemoveBooks(ctx, db, book)
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
func BorrowHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		session, _ := session.Store.Get(r, "very-secret-key")
		userID, _ := session.Values[library.SessionKeyUserId].(string)

		book_id, _ := strconv.Atoi(r.FormValue("book_id"))
		book := library.Borrowed_Book{
			User_id:     userID,
			Book_id:     book_id,
			Borrow_Date: time.Now(),
		}
		queries.BorrowBook(ctx, db, book)
		http.Redirect(w, r, "/your_books", http.StatusSeeOther)
	}
}
func AddBooksHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			server.RenderTemplate(w, "add_books.html", nil)
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

		book := library.Book{
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
