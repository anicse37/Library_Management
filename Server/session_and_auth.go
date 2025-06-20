package server

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("very-secret-key"))

func init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   1800, // 30 minutes
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false, // set to true if using HTTPS
	}
}
func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "library-session")
		_, userOk := session.Values["userid"]
		_, roleOk := session.Values["role"]

		if !userOk || !roleOk {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}
