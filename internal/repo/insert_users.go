package librarySQL

import (
	"context"

	"github.com/anicse37/Library_Management/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertSuperAdmin(ctx context.Context, db models.Database, user models.User) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if _, err := db.DB.ExecContext(ctx, `INSERT IGNORE INTO user
		VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
	}
}

func InsertUsers(ctx context.Context, db models.Database, user models.User) error {
	if user.Role == "user" {
		user.Approved = true
	} else {
		user.Approved = false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := db.DB.ExecContext(ctx, `INSERT INTO user
	VALUES (?,?,?,?,?);`, user.Name, user.Id, user.Role, hashedPassword, user.Approved); err != nil {
		return err
	}
	return nil
}
