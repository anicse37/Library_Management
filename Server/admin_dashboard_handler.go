package server

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
)

func AdminDashboard(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "library-session")

		userID, ok := session.Values["userid"].(string)
		userRole, ok2 := session.Values["role"].(string)

		if !ok || !ok2 || userID == "" || userRole != "admin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		user, err := db.GetUserByID(ctx, userID, "id")
		if err != nil || !user.Approved || user.Role != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		RenderTemplate(w, "admin_dashboard.html", nil)
	}
}

func AllUsersHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var users library.ListUser
			if search != "" {
				users = db.SearchUsers(ctx, search)
			} else {
				users, _ = library.GetAllUser(ctx, db)
			}
			data := struct {
				Users library.ListUser
				Query string
			}{
				Users: users,
				Query: search,
			}
			RenderTemplate(w, "all-users.html", data)
		default:

		}
	}
}
