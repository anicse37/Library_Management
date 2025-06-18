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
	query := fmt.Sprintf(`INSERT INTO user
	VALUES (%v,%v,%v,%v,%v);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved)
	if _, err := db.DB.ExecContext(ctx, query); err != nil {
		log.Fatalf("Error While Inserting: %v", err)
	}
}
