package server

import (
	"context"
	"net/http"

	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func UserHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")

		id := session.Values[models.SessionKeyUserId].(string)
		User, err := queries.GetAdminWithId(ctx, db, id)
		if err != nil {
			http.Redirect(w, r, "/home?msg=login_failed", http.StatusSeeOther)
		}
		data := struct {
			User models.User
		}{
			User: User,
		}
		template.RenderTemplate(w, "home_user.html", data)
	}
}
