package server

import (
	"context"
	"net/http"
	"time"

	library "github.com/anicse37/Library_Management/Files"
)

func registerHandler(ctx context.Context, db library.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		id := r.FormValue("id")
		role := r.FormValue("role")
		db.InsertUser(ctx, library.User{
			Id:       id,
			Name:     username,
			Password: password,
			Role:     role,
		})
	}
}

func ServeLibrary(dns string) {
	db := library.Database{}
	db.DB = library.StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	router := http.NewServeMux()
	router.HandleFunc("/register", registerHandler(ctx, db))

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Server/static"))))

	http.ListenAndServe(":8080", router)
}
