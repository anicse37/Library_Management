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
	db.InsertSuperAdmin(ctx, library.User{
		Id:       "MAU21UCS014",
		Name:     "Super Admin",
		Role:     "superadmin",
		Password: "Aniket@9811",
		Approved: true,
	})
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("Server/static/css"))
	router.Handle("/Server/static/css/", http.StripPrefix("/Server/static/css/", fs))

	router.HandleFunc("/register", server.RegisterHandler(ctx, db))
	router.HandleFunc("/login", server.LoginHandler(ctx, db))
	router.HandleFunc("/logout", server.LogoutHandler())

	router.HandleFunc("/books", server.RequireLogin(server.BooksHandle(ctx, db)))
	router.HandleFunc("/all-users", server.RequireLogin(server.AllUsersHandler(ctx, db)))

	router.HandleFunc("/admin/dashboard", server.RequireLogin(server.RequireRole("admin", server.AdminDashboard(ctx, db))))
	router.HandleFunc("/superadmin/dashboard", server.RequireLogin(server.RequireRole("superadmin", server.SuperAdminDashboard(ctx, db))))
	router.HandleFunc("/approve-users", server.RequireLogin(server.RequireRole("superadmin", server.ApproveUsers(ctx, db))))

	http.ListenAndServe(":5050", router)
}
