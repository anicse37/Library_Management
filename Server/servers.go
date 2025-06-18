package server

import (
	"context"
	"net/http"
	"text/template"

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

			// ✅ At this point, login is successful —> set cookies
			http.SetCookie(w, &http.Cookie{
				Name:     "username",
				Value:    user.Name,
				Path:     "/",
				HttpOnly: true,
			})

			http.SetCookie(w, &http.Cookie{
				Name:     "role",
				Value:    user.Role,
				Path:     "/",
				HttpOnly: true,
			})

			// Redirect to books page
			http.Redirect(w, r, "/books", http.StatusSeeOther)

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

			data := struct {
				Book  []library.BookJSON
				Query string
			}{
				Book:  books.Book,
				Query: search,
			}
			RenderTemplate(w, "books.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("Server/static/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
