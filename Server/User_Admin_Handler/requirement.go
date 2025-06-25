package handler

import (
	"fmt"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		_, userOk := session.Values[library.SessionKeyUserId]
		_, roleOk := session.Values[library.SessionKeyRole]

		if !userOk || !roleOk {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}
func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		if rRole, ok := session.Values[library.SessionKeyRole].(string); !ok || rRole != role {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
func RequireTwoRoles(role1 string, role2 string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "very-secret-key")
		if rRole, ok := session.Values[library.SessionKeyRole].(string); !ok || (rRole != role1 && rRole != role2) {
			fmt.Println(rRole, role1, role2)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
