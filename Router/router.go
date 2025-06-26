package router

import (
	"context"
	"net/http"
	"time"

	server "github.com/anicse37/Library_Management/internal/handlers"
	books "github.com/anicse37/Library_Management/internal/handlers/Books"
	dashboard "github.com/anicse37/Library_Management/internal/handlers/Dashboard"
	authentication "github.com/anicse37/Library_Management/internal/handlers/Login_Register"
	handler "github.com/anicse37/Library_Management/internal/handlers/User_Admin_Handler"
	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func Router(dns string, SuperAdmin models.User) {
	db := models.Database{}
	db.DB = librarySQL.StartDB(dns)
	defer db.DB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()
	// librarySQL.InsertSuperAdmin(ctx, db, SuperAdmin)
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("internal/template/static/css"))
	router.Handle("/internal/template/static/css/", http.StripPrefix("/internal/template/static/css/", fs))

	router.HandleFunc("/register", authentication.RegisterHandler(ctx, db))
	router.HandleFunc("/login", authentication.LoginHandler(ctx, db))
	router.HandleFunc("/logout", authentication.LogoutHandler())

	router.HandleFunc("/books", handler.RequireLogin(books.BooksHandle(ctx, db)))
	router.HandleFunc("/add_book", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", handler.AddBooksHandler(ctx, db))))
	router.HandleFunc("/remove_books", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", handler.RemoveBooksHandler(ctx, db))))
	router.HandleFunc("/your_books", handler.RequireLogin(handler.RequireRole("user", handler.BorrowedBooksHandler(ctx, db))))
	router.HandleFunc("/borrow", handler.RequireLogin(handler.RequireRole("user", handler.BorrowHandler(ctx, db))))

	router.HandleFunc("/home", handler.RequireLogin(handler.RequireRole("user", server.UserHandler(ctx, db))))

	router.HandleFunc("/remove_user", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", handler.RemoveUserHandler(ctx, db))))
	router.HandleFunc("/all_users", handler.RequireLogin(handler.AllUsersHandler(ctx, db)))

	router.HandleFunc("/manage_admins", handler.RequireLogin(handler.RequireRole("superadmin", handler.AllAdminsHandler(ctx, db))))
	router.HandleFunc("/approve_admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.ApproveHandler(ctx, db))))
	router.HandleFunc("/remove_admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.RemoveAdminHandler(ctx, db))))

	router.HandleFunc("/admin/dashboard", handler.RequireLogin(handler.RequireRole("admin", dashboard.AdminDashboard(ctx, db))))
	router.HandleFunc("/superadmin/dashboard", handler.RequireLogin(handler.RequireRole("superadmin", dashboard.SuperAdminDashboard(ctx, db))))
	http.ListenAndServe(":5050", router)

}
