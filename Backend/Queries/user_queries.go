package queries

import (
	"context"

	library "github.com/anicse37/Library_Management/Backend"
)

func GetAllUsers(ctx context.Context, db library.Database) (library.ListUser, error) {
	users, err := library.GetWithRoles(ctx, db, library.RoleUser)
	return users, err
}
func GetAllAdmins(ctx context.Context, db library.Database) (library.ListUser, error) {
	users, err := library.GetWithRoles(ctx, db, library.RoleAdmin)
	return users, err
}

func GetUserWithId(ctx context.Context, db library.Database, id string) (library.User, error) {
	user, err := library.GetWithID(ctx, db, id, library.RoleUser)
	return user, err
}
func GetAdminWithId(ctx context.Context, db library.Database, id string) (library.User, error) {
	user, err := library.GetWithID(ctx, db, id, library.RoleAdmin)
	return user, err
}

func GetAdminsApproved(ctx context.Context, db library.Database) (library.ListUser, error) {
	users, err := library.GetAdminsWithApprovals(ctx, db, 1)
	return users, err
}

func GetAdminsNotApproved(ctx context.Context, db library.Database) (library.ListUser, error) {
	users, err := library.GetAdminsWithApprovals(ctx, db, 0)
	return users, err
}

func SearchUsers(ctx context.Context, db library.Database, keyword string) (library.ListUser, error) {
	user, err := library.SearchWithRole(ctx, db, library.RoleUser, keyword)
	return user, err
}
func SearchAdmins(ctx context.Context, db library.Database, keyword string) (library.ListUser, error) {
	user, err := library.SearchWithRole(ctx, db, library.RoleAdmin, keyword)
	return user, err
}

func ApproveAdmin(ctx context.Context, db library.Database, id string) {
	db.DB.ExecContext(ctx, `UPDATE user SET approved = 1 WHERE id = ?`, id)
}

func RemoveAdmin(ctx context.Context, db library.Database, id string) {
	db.DB.ExecContext(ctx, `DELETE FROM user WHERE id = ?`, id)
}
func RemoveUser(ctx context.Context, db library.Database, id string) {
	db.DB.ExecContext(ctx, `DELETE FROM borrowed_books WHERE id = ?`, id)
	db.DB.ExecContext(ctx, `DELETE FROM user WHERE id = ?`, id)
}
