package library

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrorScanningUser  = errors.New("Error Scanning user from database:")
	ErrorScanningUsers = errors.New("Error Scanning users from database:")
)

func GetWithRoles(ctx context.Context, db Database, role string) (ListUser, error) {
	users := ListUser{}
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM user WHERE role = ?", role)
	if err != nil {
		return users, ErrorScanningUsers
	}
	defer result.Close()

	users = ScanUsers(result)
	return users, nil
}
func GetAdminsWithApprovals(ctx context.Context, db Database, approval string) (ListUser, error) {
	users := ListUser{}
	result, err := db.DB.QueryContext(ctx, "SELECT * FROM user WHERE role= 'admin' AND approval = '?'", approval)
	if err != nil {
		return users, ErrorScanningUsers
	}
	defer result.Close()

	users = ScanUsers(result)
	return users, nil
}

func GetWithID(ctx context.Context, db Database, id string, role string) (User, error) {
	user := User{}
	res := db.DB.QueryRowContext(ctx, `SELECT name, id, role, password, approved FROM ? WHERE id = '?';`, role, id)

	err := res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
	if err != nil {
		return user, ErrorScanningUser
	}

	return user, nil
}

// Below are the helper functions.
func ScanUsers(result *sql.Rows) ListUser {
	users := ListUser{}
	user := User{}
	for result.Next() {
		result.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
		users.Users = append(users.Users, user)
	}
	return users
}
