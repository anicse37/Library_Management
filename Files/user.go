package library

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *Database) InsertUser(ctx context.Context, user User) {
	if user.Role == "user" {
		user.Approved = true
	} else {
		user.Approved = false
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if _, err := db.DB.ExecContext(ctx, `INSERT INTO user
	VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
		fmt.Printf("Error While Inserting: %v\n", err)
	}
}
func (db *Database) InsertSuperAdmin(ctx context.Context, user User) {
	_, err := db.GetUserByID(ctx, user.Id, "id")
	if err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if _, err := db.DB.ExecContext(ctx, `INSERT INTO user
	VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
			fmt.Printf("Error While Inserting: %v\n", err)
		}
	}
}

func (db *Database) GetUserByID(ctx context.Context, name string, field string) (User, error) {
	user := User{}
	switch field {
	case "id":
		find := fmt.Sprintf(`SELECT name, id, role, password, approved FROM user WHERE id = '%v';`, name)
		res := db.DB.QueryRowContext(ctx, find)
		err := res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
		if err != nil {
			log.Printf("Error scanning user from DB: %v\n", err)
			return user, fmt.Errorf("user not found or DB error")
		}

	default:
		fmt.Println("Invalid input")
		return user, nil
	}
	return user, nil
}
func GetApprovalUsers(ctx context.Context, db Database) ListUser {
	users := ListUser{}
	user := User{}
	pendingApproval, err := db.DB.Query(`SELECT * FROM user WHERE approved = FALSE`)
	if err != nil {
		log.Println("Error querying DB:", err)
		return users
	}
	defer pendingApproval.Close()

	for pendingApproval.Next() {
		pendingApproval.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users
}

func GetAllUser(ctx context.Context, db Database) (ListUser, error) {
	users := ListUser{}
	user := User{}
	res, err := db.DB.Query(`SELECT * FROM user WHERE 	role = 'user'`)
	if err != nil {
		log.Printf("Error While Loading All Users: %v", err)
		return users, err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users, nil
}
func (db *Database) SearchUsers(ctx context.Context, keyword string) ListUser {
	users := ListUser{}
	user := User{}
	query := `SELECT * FROM user WHERE name LIKE ? OR id LIKE ?;`
	likePattern := "%" + keyword + "%"
	res, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return users
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users
}

/*working on it*/

func GetAllUsersWithRole(ctx context.Context, db Database) (ListUser, error) {
	users := ListUser{}
	user := User{}
	res, err := db.DB.Query(`SELECT * FROM user WHERE 	role = 'user'`)
	if err != nil {
		log.Printf("Error While Loading All Users: %v", err)
		return users, err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users, nil
}
