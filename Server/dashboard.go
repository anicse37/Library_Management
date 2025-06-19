package server

import (
	"context"
	"net/http"
	"text/template"

	library "github.com/anicse37/Library_Management/Files"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("Server/static/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AdminDashboard(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "library-session")
		userID, _ := session.Values["role"].(string)
		if userID != "admin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		user, _ := db.GetUserByID(ctx, userID, "role")
		if !user.Approved || user.Role != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		RenderTemplate(w, "admin_dashboard.html", nil)
	}

}
func SuperAdminDashboard(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "library-session")
		userRole, ok := session.Values["role"]
		if !ok || userRole != "superadmin" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		userID := r.FormValue("id")
		user, err := db.GetUserByID(ctx, userID, "id")
		if err != nil || !user.Approved || user.Role != "superadmin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		RenderTemplate(w, "superadmin_dashboard.html", nil)
	}
}
