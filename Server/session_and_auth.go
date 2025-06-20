package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))

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
		session, _ := Store.Get(r, "very-secret-key")
		_, userOk := session.Values["userid"]
		_, roleOk := session.Values["role"]

		if !userOk || !roleOk {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}
