package handler

import (
	"context"
	"fmt"
	"net/http"

	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
	"github.com/anicse37/Library_Management/internal/search"
	queries "github.com/anicse37/Library_Management/internal/services"
	"github.com/anicse37/Library_Management/internal/template"
)

func AllUsersHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			searchQuerry := r.URL.Query().Get("search")
			var users models.ListUser
			if searchQuerry != "" {
				users, _ = search.SearchUsers(ctx, db, "admin", searchQuerry)
			} else {
				users, _ = queries.GetAllUsers(ctx, db)
			}
			role := "user"
			session, _ := session.Store.Get(r, "very-secret-key")
			if rRole, ok := session.Values[models.SessionKeyRole].(string); ok {
				role = rRole
			}

			data := struct {
				Users models.ListUser
				Query string
				Role  string
			}{
				Users: users,
				Query: searchQuerry,
				Role:  role,
			}
			template.RenderTemplate(w, "all_users.html", data)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
func AllAdminsHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			search := r.URL.Query().Get("search")
			user1, _ := queries.GetAdminsNotApproved(ctx, db)
			user2, _ := queries.GetAdminsApproved(ctx, db)

			data := struct {
				Users1 models.ListUser
				Users2 models.ListUser
				Querry string
			}{
				Users1: user1,
				Users2: user2,
				Querry: search,
			}
			template.RenderTemplate(w, "manage_admins.html", data)
		default:
		}
	}
}

func ApproveHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		admin := r.FormValue("admin_id")
		fmt.Println(admin)
		queries.ApproveAdmin(ctx, db, admin)
		http.Redirect(w, r, "/manage_admins", http.StatusSeeOther)
	}
}
func RemoveAdminHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		admin := r.FormValue("admin_id")
		queries.RemoveAdmin(ctx, db, admin)
		http.Redirect(w, r, "/manage_admins", http.StatusSeeOther)
	}
}
func RemoveUserHandler(ctx context.Context, db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.FormValue("user_id")
		queries.RemoveUser(ctx, db, user)
		http.Redirect(w, r, "/all_users", http.StatusSeeOther)
	}
}
