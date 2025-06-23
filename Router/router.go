package router

import (
	"context"
	"net/http"
	"time"

	library "github.com/anicse37/Library_Management/Files"
	server "github.com/anicse37/Library_Management/Server"
	_ "github.com/go-sql-driver/mysql" // <-- THIS is the fix
)

func Router(dns string, SuperAdmin library.User) {
	db := library.Database{}
	db.DB = library.StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)

	defer cancel()
	db.InsertSuperAdmin(ctx, SuperAdmin)
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("Server/static/css"))
	router.Handle("/Server/static/css/", http.StripPrefix("/Server/static/css/", fs))

	router.HandleFunc("/register", server.RegisterHandler(ctx, db))
	router.HandleFunc("/login", server.LoginHandler(ctx, db))
	router.HandleFunc("/logout", server.LogoutHandler())

	router.HandleFunc("/home", server.RequireLogin(server.Home(ctx, db)))

	router.HandleFunc("/books", server.RequireLogin(server.BooksHandle(ctx, db)))
	router.HandleFunc("/all-users", server.RequireLogin(server.AllUsersHandler(ctx, db)))
	router.HandleFunc("/all-admins", server.RequireLogin(server.AllAdminsHandler(ctx, db)))

	router.HandleFunc("/admin/dashboard", server.RequireLogin(server.RequireRole("admin", server.AdminDashboard(ctx, db))))
	router.HandleFunc("/superadmin/dashboard", server.RequireLogin(server.RequireRole("superadmin", server.SuperAdminDashboard(ctx, db))))
	router.HandleFunc("/approve-users", server.RequireLogin(server.RequireRole("superadmin", server.ApproveUsers(ctx, db))))

	http.ListenAndServe(":5050", router)
}
