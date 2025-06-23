package server

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
)

func SuperAdminDashboard(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "very-secret-key")

		userID, ok := session.Values["userid"].(string)
		userRole, ok2 := session.Values["role"].(string)

		if !ok || !ok2 || userID == "" || userRole != "superadmin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		user, err := db.GetUserByID(ctx, userID, "id")
		if err != nil || !user.Approved || user.Role != "superadmin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		RenderTemplate(w, "superadmin_dashboard.html", nil)
	}
}

func ApproveUsers(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "very-secret-key")

		userID, ok := session.Values["userid"].(string)
		userRole, ok2 := session.Values["role"].(string)

		if !ok || !ok2 || userID == "" || userRole != "superadmin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case http.MethodGet:
			user := library.GetApprovalUsers(ctx, db)
			data := struct {
				User library.ListUser
			}{
				User: user,
			}
			RenderTemplate(w, "approve-user.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	}
}
