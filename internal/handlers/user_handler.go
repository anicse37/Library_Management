package server

import (
	"context"
	"fmt"
	"net/http"

	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func UserHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := session.Store.Get(r, "very-secret-key")
		if err != nil {
			fmt.Printf("error yaha hai: %v\n", err)

		}

		id := session.Values[models.SessionKeyUserId].(string)
		User, err := queries.GetAdminWithId(ctx, db, id)
		if err != nil {
			fmt.Printf("error yaha hai: %v\n", err)
		}
		data := struct {
			User models.User
		}{
			User: User,
		}
		template.RenderTemplate(w, "home_user.html", data)
	}
}
