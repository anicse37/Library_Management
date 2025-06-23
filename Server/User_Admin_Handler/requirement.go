package handler

import (
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
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
