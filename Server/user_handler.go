package server

import (
	"context"
	"fmt"
	"net/http"

	library "github.com/anicse37/Library_Management/Files"
)

func Home(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "very-secret-key")
		if err != nil {
			fmt.Printf("error yaha hai: %v\n", err)

		}

		id := session.Values["userid"].(string)
		User, err := db.GetUserByID(ctx, id, "id")
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
