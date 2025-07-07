package authentication

import (
	"context"
	"net/http"

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
				http.Redirect(w, r, "/login?msg=login_failed", http.StatusSeeOther)
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				http.Redirect(w, r, "/login?msg=login_failed", http.StatusSeeOther)
				return
			}

			if !user.Approved {
				http.Redirect(w, r, "/login?msg=admin_not_approved", http.StatusSeeOther)
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
				http.Redirect(w, r, "admin/dashboard?msg=login_success", http.StatusSeeOther)
			case "superadmin":
				http.Redirect(w, r, "superadmin/dashboard?msg=login_success", http.StatusSeeOther)
			default:
				http.Redirect(w, r, "home?msg=login_success", http.StatusSeeOther)
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

		http.Redirect(w, r, "/login?msg=logout_success", http.StatusSeeOther)
	}
}
