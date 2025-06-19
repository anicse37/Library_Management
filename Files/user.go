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
