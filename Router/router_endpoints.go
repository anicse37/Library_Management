package router

import (
	"context"
	"net/http"

	server "github.com/anicse37/Library_Management/internal/handlers"
	books "github.com/anicse37/Library_Management/internal/handlers/Books"
	dashboard "github.com/anicse37/Library_Management/internal/handlers/Dashboard"
	authentication "github.com/anicse37/Library_Management/internal/handlers/Login_Register"
	handler "github.com/anicse37/Library_Management/internal/handlers/User_Admin_Handler"
	"github.com/anicse37/Library_Management/internal/models"
)

func RouterEndpoints(ctx context.Context, db models.Database) *http.ServeMux {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("internal/template/static/css"))
	router.Handle("/internal/template/static/css/", http.StripPrefix("/internal/template/static/css/", fs))

	router.HandleFunc("/register", authentication.RegisterHandler(ctx, db))
	router.HandleFunc("/login", authentication.LoginHandler(ctx, db))
	router.HandleFunc("/logout", authentication.LogoutHandler())

	router.HandleFunc("/books", handler.RequireLogin(books.BooksHandle(ctx, db)))
	router.HandleFunc("/add_book", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", books.AddBooksHandler(ctx, db))))
	router.HandleFunc("/remove_books", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", books.RemoveBooksHandler(ctx, db))))
	router.HandleFunc("/your_books", handler.RequireLogin(handler.RequireRole("user", books.BorrowedBooksHandler(ctx, db))))
	router.HandleFunc("/borrow", handler.RequireLogin(handler.RequireRole("user", books.BorrowHandler(ctx, db))))
	router.HandleFunc("/return_book", handler.RequireLogin(handler.RequireRole("user", books.ReturnBookHandler(ctx, db))))

	router.HandleFunc("/home", handler.RequireLogin(handler.RequireRole("user", server.UserHandler(ctx, db))))

	router.HandleFunc("/remove_user", handler.RequireLogin(handler.RequireTwoRoles("admin", "superadmin", handler.RemoveUserHandler(ctx, db))))
	router.HandleFunc("/all_users", handler.RequireLogin(handler.AllUsersHandler(ctx, db)))

	router.HandleFunc("/manage_admins", handler.RequireLogin(handler.RequireRole("superadmin", handler.AllAdminsHandler(ctx, db))))
	router.HandleFunc("/approve_admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.ApproveHandler(ctx, db))))
	router.HandleFunc("/remove_admin", handler.RequireLogin(handler.RequireRole("superadmin", handler.RemoveAdminHandler(ctx, db))))

	router.HandleFunc("/admin/dashboard", handler.RequireLogin(handler.RequireRole("admin", dashboard.AdminDashboard(ctx, db))))
	router.HandleFunc("/superadmin/dashboard", handler.RequireLogin(handler.RequireRole("superadmin", dashboard.SuperAdminDashboard(ctx, db))))

	router.HandleFunc("/error", handler.ErrorHandler(ctx, db))

	return router
}
