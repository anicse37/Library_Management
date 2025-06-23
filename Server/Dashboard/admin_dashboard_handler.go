package dashboard

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func AdminDashboard(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")

		userID, ok := session.Values[library.SessionKeyUserId].(string)
		userRole, ok2 := session.Values[library.SessionKeyRole].(string)

		if !ok || !ok2 || userID == "" || userRole != "admin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		user, err := db.GetUserByID(ctx, userID, library.SessionKeyUserId)
		if err != nil || !user.Approved || user.Role != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		server.RenderTemplate(w, "admin_dashboard.html", nil)
	}
}
