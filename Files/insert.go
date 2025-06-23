package library

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *Database) InsertSuperAdmin(ctx context.Context, user User) {
	_, err := db.GetUserByID(ctx, user.Id, SessionKeyUserId)
	if err != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if _, err := db.DB.ExecContext(ctx, `INSERT INTO user
	VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
			fmt.Printf("Error While Inserting: %v\n", err)
		}
	}
}

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
func (db *Database) InsertBooksInTable(ctx context.Context, book Book) {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description,year)
	VALUES (?,?,?,?);`, book.Name, book.Author, book.Description, book.Year); err != nil {
		log.Fatalf("Error While Inserting: %v", err)
	}
}
