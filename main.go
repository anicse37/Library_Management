package main

import (
	"context"
	"net/http"
	"time"

	library "github.com/anicse37/Library_Management/Files"
	server "github.com/anicse37/Library_Management/Server"
	_ "github.com/go-sql-driver/mysql" // <-- THIS is the fix
)

const dns = "go_user:S3cur3P@ssw0rd@tcp(127.0.0.1:3306)/library_db?charset=utf8mb4&parseTime=true&loc=Local"

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func main() {
	db := library.Database{}
	db.DB = library.StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	router := http.NewServeMux()
	router.HandleFunc("/register", server.RegisterHandler(ctx, db))
	router.HandleFunc("/login", server.LoginHandler(ctx, db))
	router.HandleFunc("/books", server.BooksHandle(ctx, db))

	http.ListenAndServe(":8080", router)
}
