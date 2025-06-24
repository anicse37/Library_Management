package router

import (
	"context"
	"net/http"
	"time"

	library "github.com/anicse37/Library_Management/Backend"
	queries "github.com/anicse37/Library_Management/Backend/Queries"
	server "github.com/anicse37/Library_Management/Server"
	books "github.com/anicse37/Library_Management/Server/Books"
	dashboard "github.com/anicse37/Library_Management/Server/Dashboard"
	authentication "github.com/anicse37/Library_Management/Server/Login_Register"
	handler "github.com/anicse37/Library_Management/Server/User_Admin_Handler"
)

func Router(dns string, SuperAdmin library.User) {
	db := library.Database{}
	db.DB = StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()
	queries.InsertSuperAdmin(ctx, db, SuperAdmin)
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("Server/static/css"))
	router.Handle("/Server/static/css/", http.StripPrefix("/Server/static/css/", fs))

	router.HandleFunc("/register", authentication.RegisterHandler(ctx, db))
	router.HandleFunc("/login", authentication.LoginHandler(ctx, db))
	router.HandleFunc("/logout", authentication.LogoutHandler())

	router.HandleFunc("/home", handler.RequireLogin(server.UserHandler(ctx, db)))

	router.HandleFunc("/books", handler.RequireLogin(books.BooksHandle(ctx, db)))
	router.HandleFunc("/borrowed-books", handler.RequireLogin(books.BorrowedBooksHandle(ctx, db)))

	router.HandleFunc("/all-users", handler.RequireLogin(handler.AllUsersHandler(ctx, db)))
	router.HandleFunc("/manage-admins", handler.RequireLogin(handler.RequireRole("superadmin", handler.AllAdminsHandler(ctx, db))))
	router.HandleFunc("/approve-admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.ApproveHandler(ctx, db))))
	router.HandleFunc("/remove-admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.RemoveHandler(ctx, db))))

	router.HandleFunc("/admin/dashboard", handler.RequireLogin(handler.RequireRole("admin", dashboard.AdminDashboard(ctx, db))))
	router.HandleFunc("/superadmin/dashboard", handler.RequireLogin(handler.RequireRole("superadmin", dashboard.SuperAdminDashboard(ctx, db))))
	http.ListenAndServe(":5050", router)
}
