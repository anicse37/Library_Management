package server

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	library "github.com/anicse37/Library_Management/Files"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Parse form and insert user
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form submission", http.StatusBadRequest)
				return
			}
			username := r.FormValue("username")
			password := r.FormValue("password")
			id := r.FormValue("id")
			role := r.FormValue("role")

			if username == "" || password == "" || id == "" || role == "" {
				http.Error(w, "Missing form fields", http.StatusBadRequest)
				return
			}

			db.InsertUser(ctx, library.User{
				Id:       id,
				Name:     username,
				Password: password,
				Role:     role,
			})
			http.Redirect(w, r, "/register", http.StatusSeeOther) // Redirect after successful POST
			return
		}

		// For GET requests, just render the form
		RenderTemplate(w, "register.html", nil)
	}
}
func LoginHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form submission", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			password := r.FormValue("password")

			if id == "" || password == "" {
				http.Error(w, "Missing form fields", http.StatusBadRequest)
				return
			}

			user := db.FindUser(ctx, library.User{
				Id:       id,
				Password: password,
			}, "id")
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				fmt.Printf("Invalid Password or username: %v", err)
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		RenderTemplate(w, "login.html", nil)
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
