package queries

import (
	"context"

	library "github.com/anicse37/Library_Management/Backend"
	"github.com/anicse37/Library_Management/internal/models"
	librarySQL "github.com/anicse37/Library_Management/internal/repo"
)

func GetAllUsers(ctx context.Context, db models.Database) (models.ListUser, error) {
	users, err := librarySQL.GetWithRoles(ctx, db, models.RoleUser)
	return users, err
}
func GetAllAdmins(ctx context.Context, db models.Database) (models.ListUser, error) {
	users, err := librarySQL.GetWithRoles(ctx, db, models.RoleAdmin)
	return users, err
}

func GetUserWithId(ctx context.Context, db models.Database, id string) (models.User, error) {
	user, err := librarySQL.GetWithID(ctx, db, id, models.RoleUser)
	return user, err
}
func GetAdminWithId(ctx context.Context, db models.Database, id string) (models.User, error) {
	user, err := librarySQL.GetWithID(ctx, db, id, models.RoleAdmin)
	return user, err
}

func GetAdminsApproved(ctx context.Context, db models.Database) (models.ListUser, error) {
	users, err := librarySQL.GetAdminsWithApprovals(ctx, db, 1)
	return users, err
}

func GetAdminsNotApproved(ctx context.Context, db models.Database) (models.ListUser, error) {
	users, err := librarySQL.GetAdminsWithApprovals(ctx, db, 0)
	return users, err
}

func SearchUsers(ctx context.Context, db models.Database, keyword string) (models.ListUser, error) {
	user, err := library.SearchWithRole(ctx, db, models.RoleUser, keyword)
	return user, err
}
func SearchAdmins(ctx context.Context, db models.Database, keyword string) (models.ListUser, error) {
	user, err := library.SearchWithRole(ctx, db, models.RoleAdmin, keyword)
	return user, err
}

func ApproveAdmin(ctx context.Context, db models.Database, id string) {
	db.DB.ExecContext(ctx, `UPDATE user SET approved = 1 WHERE id = ?`, id)
}

func RemoveAdmin(ctx context.Context, db models.Database, id string) {
	db.DB.ExecContext(ctx, `DELETE FROM user WHERE id = ?`, id)
}
func RemoveUser(ctx context.Context, db models.Database, id string) {
	db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE id = ?`, id)
	db.DB.ExecContext(ctx, `DELETE FROM user WHERE id = ?`, id)
}
