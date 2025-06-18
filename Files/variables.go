package library

import (
	"context"
	"database/sql"
)

type Library interface {
	GetBooksFromTable(ctx context.Context)
	InsertInTable(ctx context.Context, book ListBookJSON)
	InsertUser(ctx context.Context, user User)
	Display(name string)
}
type Database struct {
	Files Library
	DB    *sql.DB
}
type ListBookJSON struct {
	Book []BookJSON `json:"book"`
}
type BookJSON struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	Description string `json:"description"`
	Available   int    `json:"available"`
}
type Book struct {
	Name        string
	Author      string
	Year        int
	Description string
}
type User struct {
	Name     string
	Id       string
	Password string
	Role     string
	Approved bool
}

const (
	ColunmTypeINT          = "INT"
	ColunmTypeVARCHAR      = "VARCHAR(255)"
	ColunmTypeBOOLEAN      = "BOOLEAN"
	ColunmTypeBOOLEANTrue  = "BOOLEAN DEFAULT TRUE"
	ColunmTypeBOOLEANFalse = "BOOLEAN DEFAULT FALSE"
)
