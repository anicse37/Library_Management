package dashboard

import (
	"context"
	"net/http"

	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func AdminDashboard(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		userID, ok := session.Values[models.SessionKeyUserId].(string)
		userRole, ok2 := session.Values[models.SessionKeyRole].(string)
		if !ok || !ok2 || userID == "" || userRole != "admin" {
			http.Redirect(w, r, "/login?msg=unauthorized_access", http.StatusUnauthorized)
			return
		}

		user, err := queries.GetAdminWithId(ctx, db, userID)
		if err != nil || !user.Approved || user.Role != "admin" {
			http.Redirect(w, r, "/login?msg=unauthorized_access", http.StatusUnauthorized)
			return
		}
		data := struct {
			Name string
		}{
			Name: user.Name,
		}

		template.RenderTemplate(w, "admin_dashboard.html", data)
	}
}
