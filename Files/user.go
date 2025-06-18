package library

import (
	"context"
	"fmt"

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

func (db *Database) FindUser(ctx context.Context, name any, field string) User {
	var find string
	user := User{}
	switch field {
	case "id":
		find = fmt.Sprintf(`SELECT * FROM user
		WHERE %v = ?;`, field)
	default:
		fmt.Println("Invalid input")
	}
	res := db.DB.QueryRowContext(ctx, find, name)
	err := res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
	if err != nil {
		fmt.Println("User Doesn't extst:")
	}
	return user
}
