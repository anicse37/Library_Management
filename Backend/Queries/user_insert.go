package queries

import (
	"context"
	"fmt"

	library "github.com/anicse37/Library_Management/Backend"
	"golang.org/x/crypto/bcrypt"
)

func InsertSuperAdmin(ctx context.Context, db library.Database, user library.User) {
	_, err := GetUserWithId(ctx, db, user.Id)
	if err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if _, err := db.DB.ExecContext(ctx, `INSERT INTO user
		VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
			fmt.Printf("Error While Inserting: %v\n", err)
		}
	}
}

func InsertUsers(ctx context.Context, db library.Database, user library.User) {
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
