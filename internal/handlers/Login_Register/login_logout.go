package authentication

import (
	"context"
	"net/http"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			r.ParseForm()
			id := r.FormValue(models.SessionKeyUserId)
			password := r.FormValue(models.SessionKeyPassword)
			user, err := queries.GetUserWithId(ctx, db, id)
			if err != nil {
				errors_package.SetError(errors_package.ErrorInvalidUser)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				errors_package.SetError(errors_package.ErrorInvalidPassword)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}

			if !user.Approved {
				errors_package.SetError(errors_package.ErrorAdminNotAllowed)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}

			session, _ := session.Store.Get(r, "very-secret-key")
			session.Values[models.SessionKeyUsername] = user.Name
			session.Values[models.SessionKeyUserId] = user.Id
			session.Values[models.SessionKeyRole] = user.Role
			session.Values[models.SessionKeyPassword] = user.Password
			session.Save(r, w)

			switch user.Role {
			case "admin":
				http.Redirect(w, r, "admin/dashboard", http.StatusSeeOther)
			case "superadmin":
				http.Redirect(w, r, "superadmin/dashboard", http.StatusSeeOther)
			default:
				http.Redirect(w, r, "home", http.StatusSeeOther)
			}
		default:
			http.ServeFile(w, r, "internal/template/static/login.html")
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
