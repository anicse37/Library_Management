package handler

import (
	"context"
	"net/http"
	"time"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func AllUsersHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var users library.ListUser
			if search != "" {
				users, _ = queries.SearchUsers(ctx, db, search)
			} else {
				users, _ = queries.GetAllUsers(ctx, db)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[library.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Users library.ListUser
				Query string
				Role  string
			}{
				Users: users,
				Query: search,
				Role:  role,
			}
			server.RenderTemplate(w, "all-users.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func AllAdminsHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			user1, _ := queries.GetAdminsNotApproved(ctx, db)
			user2, _ := queries.GetAdminsApproved(ctx, db)

			data := struct {
				Users1 library.ListUser
				Users2 library.ListUser
				Querry string
			}{
				Users1: user1,
				Users2: user2,
				Querry: search,
			}
			server.RenderTemplate(w, "manage_admins.html", data)
		default:
		}
	}
}

func ApproveHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		admin := r.FormValue("admin_id")
		queries.ApproveAdmin(ctx, db, admin)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/manage-admins", http.StatusSeeOther)
	}
}
