package authentication

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	session "github.com/anicse37/Library_Management/Server/Session"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			id := r.FormValue(library.SessionKeyUserId)
			password := r.FormValue(library.SessionKeyPassword)

			user, err := queries.GetUserWithId(ctx, db, id)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusUnauthorized)
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				return
			}

			if !user.Approved {
				http.Error(w, "Account not approved by admin", http.StatusForbidden)
				return
			}

			session, _ := session.Store.Get(r, "very-secret-key")
			session.Values[library.SessionKeyUsername] = user.Name
			session.Values[library.SessionKeyUserId] = user.Id
			session.Values[library.SessionKeyRole] = user.Role
			session.Values[library.SessionKeyPassword] = user.Password
			session.Save(r, w)

			switch user.Role {
			case "admin":
				http.Redirect(w, r, "admin/dashboard", http.StatusSeeOther)
			case "superadmin":
				http.Redirect(w, r, "superadmin/dashboard", http.StatusSeeOther)
			default:
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		default:
			http.ServeFile(w, r, "Server/static/login.html")
		}
	}
}
func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
