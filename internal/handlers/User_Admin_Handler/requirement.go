package handler

import (
	"net/http"

	errors_package "github.com/anicse37/Library_Management/internal/errors"
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
			errors_package.SetError(errors_package.ErrorUnauthorized)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
func RequireTwoRoles(role1 string, role2 string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		if rRole, ok := session.Values[models.SessionKeyRole].(string); !ok || (rRole != role1 && rRole != role2) {
			errors_package.SetError(errors_package.ErrorUnauthorized)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
