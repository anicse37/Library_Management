package models

import (
	"database/sql"
	"time"
)

type Database struct {
	DB *sql.DB
}
type Book struct {
	Id          int
	Name        string
	Author      string
	Year        int
	Description string
	Available   int
}
type ListBooks []Book

type User struct {
	Name     string
	Id       string
	Password string
	Role     string
	Approved bool
}
type ListUser []User

type Borrowed_Book struct {
	Id            int
	User_id       string
	Book_id       int
	Borrow_Date   time.Time
	Returned_Date time.Time
}
type ListBorrowed_Books []Borrowed_Book

type BorrowedBookDisplay struct {
	BorrowID     int
	UserId       string
	BookID       int
	BookName     string
	Author       string
	BorrowDate   time.Time
	ReturnedDate *time.Time
}
type ListBorrowedBookDisplay []BorrowedBookDisplay

const (
	SessionKeyUsername = "username"
	SessionKeyUserId   = "userid"
	SessionKeyRole     = "role"
	SessionKeyPassword = "password"
)
const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)
