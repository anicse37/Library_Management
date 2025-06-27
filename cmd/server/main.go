package main

import (
	"log"
	"os"

	router "github.com/anicse37/Library_Management/Router"
	"github.com/anicse37/Library_Management/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := os.Getenv("DNS")

	SuperAdmin := models.User{
		Id:       "MAU21UCS014",
		Name:     "Super Admin",
		Role:     "superadmin",
		Password: "Aniket@9811",
		Approved: true,
	}
	router.Router(dns, SuperAdmin)
}
