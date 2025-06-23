package session

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
		MaxAge:   1800,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
	}
}
