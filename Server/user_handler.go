package server

import (
	"context"
	"fmt"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
	session "github.com/anicse37/Library_Management/Server/Session"
)

func Home(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := session.Store.Get(r, "very-secret-key")
		if err != nil {
			fmt.Printf("error yaha hai: %v\n", err)

		}

		id := session.Values[library.SessionKeyUserId].(string)
		User, err := db.GetUserByID(ctx, id, library.SessionKeyUserId)
		if err != nil {
			fmt.Printf("error yaha hai: %v\n", err)
		}
		data := struct {
			User library.User
		}{
			User: User,
		}
		RenderTemplate(w, "home_user.html", data)
	}

}
