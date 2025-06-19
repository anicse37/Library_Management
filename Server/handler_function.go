package server

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			RenderTemplate(w, "register.html", nil)
			return
		}

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		id := r.FormValue("id")
		role := r.FormValue("role")

		db.InsertUser(ctx, library.User{
			Id:       id,
			Name:     username,
			Password: password,
			Role:     role,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "role",
			Value:    role,
			Path:     "/",
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    username,
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
func LoginHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			id := r.FormValue("id")
			password := r.FormValue("password")

			user, err := db.GetUserByID(ctx, id, "id")
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusUnauthorized)
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				return
			}

			if !user.Approved {
				http.Error(w, "Account not approved by admin", http.StatusForbidden)
				return
			}

			session, _ := Store.Get(r, "library-session")
			session.Values["username"] = user.Name
			session.Values["role"] = user.Role
			session.Values["password"] = user.Password
			session.Save(r, w)

			switch user.Role {
			case "admin":
				http.Redirect(w, r, "admin/dashboard", http.StatusSeeOther)
			case "superadmin":
				http.Redirect(w, r, "superadmin/dashboard", http.StatusSeeOther)
			default:
				http.Redirect(w, r, "books", http.StatusSeeOther)
			}

		default:
			http.ServeFile(w, r, "Server/static/login.html")
		}
	}
}

func BooksHandle(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var books library.ListBookJSON
			if search != "" {
				books = db.SearchBooks(ctx, search)
			} else {
				books = db.GetBooksFromTable(ctx)
			}
			role := "user"
			if roleCookier, err := r.Cookie("role"); err == nil {
				role = roleCookier.Value
			}
			data := struct {
				Book  []library.BookJSON
				Query string
				Role  string
			}{
				Book:  books.Book,
				Query: search,
				Role:  role,
			}
			RenderTemplate(w, "books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "library-session")
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
