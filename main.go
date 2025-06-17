package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	files "github.com/anicse37/SQL/Files"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = "go_user:S3cur3P@ssw0rd@tcp(127.0.0.1:3306)/library_db?charset=utf8mb4&parseTime=true&loc=Local"

func LibraryServer(w http.ResponseWriter, r *http.Request) {
	db := files.DataBase{}
	db.DB = files.StartDB(dsn)
	defer db.DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	db.CreateTable(ctx)

	switch r.Method {
	case http.MethodGet:
		data := db.Display()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	case http.MethodPost:
		var book files.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		db.InsertInTable(ctx, book)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Book added successfully"})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/books", LibraryServer)
	router.Handle("/", http.FileServer(http.Dir("Statics/")))
	http.ListenAndServe(":8080", router)
}
