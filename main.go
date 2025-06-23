package main

import (
	library "github.com/anicse37/Library_Management/Files"
	router "github.com/anicse37/Library_Management/Router"
	_ "github.com/go-sql-driver/mysql"
)

const dns = "go_user:S3cur3P@ssw0rd@tcp(127.0.0.1:3306)/library_db?charset=utf8mb4&parseTime=true&loc=Local"

func main() {
	SuperAdmin := library.User{
		Id:       "MAU21UCS014",
		Name:     "Super Admin",
		Role:     "superadmin",
		Password: "Aniket@9811",
		Approved: true,
	}
	router.Router(dns, SuperAdmin)

}
