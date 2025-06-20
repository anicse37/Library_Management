package library

import (
	"context"
	"fmt"
	"log"
)

func (db *Database) FindInTable(ctx context.Context, name any, field string) BookJSON {
	var find string
	book := BookJSON{}
	switch field {
	case "name", "author":
		find = fmt.Sprintf(`SELECT * FROM books
		WHERE %v = ?;`, field)
	default:
		fmt.Println("Invalid input")
	}
	res := db.DB.QueryRowContext(ctx, find, name)
	err := res.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
	if err != nil {
		fmt.Println("Book Doesn't extst:")
	}
	return book
}
func (db *Database) InsertBooksInTable(ctx context.Context, book Book) {
	if _, err := db.DB.ExecContext(ctx, `INSERT INTO books (name, author, description,year)
	VALUES (?,?,?,?);`, book.Name, book.Author, book.Description, book.Year); err != nil {
		log.Fatalf("Error While Inserting: %v", err)
	}
}

func (db *Database) GetBooksFromTable(ctx context.Context) ListBookJSON {
	books := ListBookJSON{}
	book := BookJSON{}
	AllBooks, err := db.DB.Query("SELECT *FROM books;")
	if err != nil {
		log.Fatalf("Error Getting Data: %v", err)
	}
	for AllBooks.Next() {
		AllBooks.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books.Book = append(books.Book, book)
	}
	return books
}
func (db *Database) SearchBooks(ctx context.Context, keyword string) ListBookJSON {
	books := ListBookJSON{}
	book := BookJSON{}

	query := `SELECT * FROM books WHERE name LIKE ? OR author LIKE ?;`
	likePattern := "%" + keyword + "%"
	rows, err := db.DB.QueryContext(ctx, query, likePattern, likePattern)
	if err != nil {
		log.Printf("Error while searching: %v", err)
		return books
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year, &book.Description, &book.Available)
		books.Book = append(books.Book, book)
	}
	return books
}
