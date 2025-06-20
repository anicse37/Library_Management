package library

import (
	"context"
	"database/sql"
)

type Library interface {
	GetBooksFromTable(ctx context.Context)
	InsertBooksInTable(ctx context.Context, book ListBookJSON)
	InsertUser(ctx context.Context, user User)
	Display(name string)
}
type Database struct {
	Files Library
	DB    *sql.DB
}
type ListBookJSON struct {
	Book []BookJSON
}
type BookJSON struct {
	Id          int
	Name        string
	Author      string
	Year        int
	Description string
	Available   int
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
type ListUser struct {
	Users []User
}
