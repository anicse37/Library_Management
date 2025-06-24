package authentication

import (
	"context"
	"net/http"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func RegisterHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			server.RenderTemplate(w, "register.html", nil)
			return
		}

		r.ParseForm()
		username := r.FormValue(library.SessionKeyUsername)
		password := r.FormValue(library.SessionKeyUsername)
		id := r.FormValue(library.SessionKeyUserId)
		role := r.FormValue(library.SessionKeyRole)

		queries.InsertUsers(ctx, db, library.User{
			Id:       id,
			Name:     username,
			Password: password,
			Role:     role,
		})

		session, _ := session.Store.Get(r, "very-secret-key")
		session.Values[library.SessionKeyUsername] = username
		session.Values[library.SessionKeyRole] = role
		session.Values[library.SessionKeyUserId] = id
		session.Save(r, w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
