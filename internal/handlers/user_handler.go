package server

import (
	"context"
	"net/http"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func UserHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := session.Store.Get(r, "very-secret-key")
		if err != nil {
			errors_package.SetError(err)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}

		id := session.Values[models.SessionKeyUserId].(string)
		User, err := queries.GetAdminWithId(ctx, db, id)
		if err != nil {
			errors_package.SetError(err)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
		data := struct {
			User models.User
		}{
			User: User,
		}
		template.RenderTemplate(w, "home_user.html", data)
	}
}
