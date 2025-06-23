package library

import (
	"context"
	"fmt"
	"log"
)

func GetAllUser(ctx context.Context, db Database) (ListUser, error) {
	users := ListUser{}
	user := User{}

	res, err := db.DB.Query(`SELECT * FROM user WHERE 	role = 'user'`)
	if err != nil {
		log.Printf("Error While Loading All Users: %v", err)
		return users, err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func GetApprovalUsers(ctx context.Context, db Database) ListUser {
	users := ListUser{}
	user := User{}
	pendingApproval, err := db.DB.Query(`SELECT * FROM user WHERE approved = FALSE`)
	if err != nil {
		log.Println("Error querying DB:", err)
		return users
	}
	defer pendingApproval.Close()

	for pendingApproval.Next() {
		pendingApproval.Scan(&user.Name, &user.Id, &user.Role, &user.Password, user.Approved)
		users.Users = append(users.Users, user)
	}
	return users
}
func (db *Database) GetUserByID(ctx context.Context, name string, field string) (User, error) {
	user := User{}
	switch field {
	case "userid":
		find := fmt.Sprintf(`SELECT name, id, role, password, approved FROM user WHERE id = '%v';`, name)
		res := db.DB.QueryRowContext(ctx, find)
		err := res.Scan(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)
		fmt.Println(&user.Name, &user.Id, &user.Role, &user.Password, &user.Approved)

		if err != nil {
			log.Printf("Error scanning user from DB: %v\n", err)
			return user, fmt.Errorf("user not found or DB error")
		}

	default:
		fmt.Println("Inalid input")
		return user, nil
	}
	return user, nil
}

func (db *Database) GetAllBooks(ctx context.Context) ListBookJSON {
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
