package handler

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
)

func AllUsersHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var users library.ListUser
			if search != "" {
				users = db.SearchUsers(ctx, search)
			} else {
				users, _ = queries.GetAllUsers(ctx, db)
			}
			data := struct {
				Users library.ListUser
				Query string
			}{
				Users: users,
				Query: search,
			}
			server.RenderTemplate(w, "all-users.html", data)
		default:
		}

	}
}
func AllAdminsHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			var users library.ListUser
			if search != "" {
				users = db.SearchUsers(ctx, search)
			} else {
				users, _ = queries.GetAllUsers(ctx, db)
			}
			data := struct {
				Users library.ListUser
				Query string
			}{
				Users: users,
				Query: search,
			}
			server.RenderTemplate(w, "all-users.html", data)
		default:
		}
	}
}
