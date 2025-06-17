package files

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Database interface {
	CreateTable(ctx context.Context, name string)
	InsertInTable(ctx context.Context, book []string)
	Display(name string)
}
type DataBase struct {
	Files Database
	DB    *sql.DB
}
type Book struct {
	Name        string
	Author      string
	Description string
	Year        int
}
type BookJSONArr struct {
	Book []BookJSON `json:"book"`
}
type BookJSON struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Available   bool   `json:"available"`
}

const (
	ColunmTypeINT          = "INT"
	ColunmTypeVARCHAR      = "VARCHAR(255)"
	ColunmTypeBOOLEAN      = "BOOLEAN"
	ColunmTypeBOOLEANTrue  = "BOOLEAN DEFAULT TRUE"
	ColunmTypeBOOLEANFalse = "BOOLEAN DEFAULT FALSE"
)

func (db *DataBase) InsertInTable(ctx context.Context, book Book) {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description, year)
	VALUES (?,?,?,?)`, book.Name, book.Author, book.Description, book.Year); err != nil {
		log.Fatalf("Inserting values Failed: %s\n", err)
	}
}

func (db *DataBase) CreateTable(ctx context.Context) {
	table := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS books (
	id %v AUTO_INCREMENT PRIMARY KEY,
	name %v NOT NULL,
	author %v,
	description %v,
	year %v,
	available %v
	);`, ColunmTypeINT, ColunmTypeVARCHAR, ColunmTypeVARCHAR, ColunmTypeVARCHAR, ColunmTypeINT, ColunmTypeBOOLEANTrue)

	if _, err := db.DB.ExecContext(ctx, table); err != nil {
		log.Fatalf("Create Table Failed: %s\n", err)
	}
}

func (Db *DataBase) Display() BookJSONArr {
	temp := `SELECT * FROM books;`
	res, err := Db.DB.Query(temp)
	if err != nil {
		log.Fatalf("Error displaying result: %v \n", err)
	}

	books := BookJSONArr{}
	for res.Next() {
		var id int
		var name, author, description string
		var year int
		var available bool

		err := res.Scan(&id, &name, &author, &description, &year, &available)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		book := BookJSON{
			Id: id, Name: name, Author: author, Description: description, Year: year, Available: available,
		}
		books.Book = append(books.Book, book)
	}
	return books
}
