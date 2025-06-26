package librarySQL

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func StartDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error Opening DB: %v\n", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping failed:v %v \n", err)
	}
	fmt.Println("Sucessfully connectedto MySQL!")

	return db
}
