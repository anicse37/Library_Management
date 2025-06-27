package librarySQL

import (
	"context"
	"database/sql"

	"github.com/anicse37/Library_Management/internal/models"
)

func GetWithRoles(ctx context.Context, db models.Database, role string) (models.ListUser, error) {
	users := &models.ListUser{}
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM user WHERE role = ?", role)
	if err != nil {
		return *users, models.ErrorScanningUsers
	}
	defer result.Close()

	users = ScanUsers(result)
	return *users, nil
}
func GetAdminsWithApprovals(ctx context.Context, db models.Database, approval int) (models.ListUser, error) {
	users := &models.ListUser{}
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM user WHERE (role= 'admin') AND (approved = ?);", approval)
	if err != nil {
		return *users, models.ErrorScanningUsers
	}
	defer result.Close()

	users = ScanUsers(result)
	return *users, nil
}

func GetWithID(ctx context.Context, db models.Database, id string, role string) (models.User, error) {
	user := models.User{}
	res := db.DB.QueryRowContext(ctx, `SELECT name, id, role, password, approved FROM user WHERE id = ?;`, id)

	err := res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
	if err != nil {
		return user, models.ErrorScanningUser
	}

	return user, nil
}

// Below are the helper functions.
func ScanUsers(result *sql.Rows) *models.ListUser {
	users := models.ListUser{}
	user := models.User{}
	for result.Next() {
		result.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
		users = append(users, user)
	}

	return &users
}
