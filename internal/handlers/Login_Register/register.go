package authentication

import (
	"context"
	"net/http"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
	"github.com/anicse37/Library_Management/internal/template"
)

func RegisterHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			template.RenderTemplate(w, "register.html", nil)
			return
		}

		r.ParseForm()
		username := r.FormValue(models.SessionKeyUsername)
		password := r.FormValue(models.SessionKeyPassword)
		id := r.FormValue(models.SessionKeyUserId)
		role := r.FormValue(models.SessionKeyRole)

		err := librarySQL.InsertUsers(ctx, db, models.User{
			Id:       id,
			Name:     username,
			Password: password,
			Role:     role,
		})
		if err != nil {
			errors_package.SetError(errors_package.ErrorUserAlreadyExist)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}

		session, _ := session.Store.Get(r, "very-secret-key")
		session.Values[models.SessionKeyUsername] = username
		session.Values[models.SessionKeyRole] = role
		session.Values[models.SessionKeyUserId] = id
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
