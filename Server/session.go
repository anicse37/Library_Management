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
