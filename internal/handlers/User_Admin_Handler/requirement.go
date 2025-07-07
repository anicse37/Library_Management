package handler

import (
	"net/http"

	session "github.com/anicse37/Library_Management/internal/middleware"
	"github.com/anicse37/Library_Management/internal/models"
)

func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		_, userOk := session.Values[models.SessionKeyUserId]
		_, roleOk := session.Values[models.SessionKeyRole]

		if !userOk || !roleOk {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}
func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		if rRole, ok := session.Values[models.SessionKeyRole].(string); !ok || rRole != role {
			if rRole == "admin" {
				http.Redirect(w, r, "/admin/dashboard?msg=unauthorized_access", http.StatusUnauthorized)
			} else if rRole == "superadmin" {
				http.Redirect(w, r, "/superadmin/dashboard?msg=unauthorized_access", http.StatusUnauthorized)
			} else {
				http.Redirect(w, r, "/home?msg=unauthorized_access", http.StatusUnauthorized)
			}
			return
		}
		next(w, r)
	}
}
func RequireTwoRoles(role1 string, role2 string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		if rRole, ok := session.Values[models.SessionKeyRole].(string); !ok || (rRole != role1 && rRole != role2) {
			http.Redirect(w, r, "/logout?msg=unauthorized_access", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
