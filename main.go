package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql" // <-- THIS is the fix

	server "github.com/anicse37/Library_Management/Server"
)

const dns = "go_user:S3cur3P@ssw0rd@tcp(127.0.0.1:3306)/library_db?charset=utf8mb4&parseTime=true&loc=Local"

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func main() {
	server.ServeLibrary(dns)
	http.ListenAndServe(":8080", nil)
}
