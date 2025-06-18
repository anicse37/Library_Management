# 📚 Library Management System (Golang + MySQL + HTML/CSS/JS)

This is a basic **Library Management System** built using:

- ⚙️ **Golang** (with `net/http` for server and `database/sql` for DB interaction)
- 🐬 **MySQL** as the database
- 🖥️ **HTML/CSS/JavaScript** for a simple frontend

---

## 🚀 Features

- Add new books via a form (name, author, description, year)
- Store book entries in a MySQL database
- Retrieve and display the book list dynamically
- JSON API at `/books` for GET and POST operations

---

## 📁 Project Structure

project-root/
│
├── main.go # Server entry point
├── Files/
│ ├── db.go # DB initialization
│ ├── books.go # DB logic: insert, display, create table
│
├── Statics/
│ ├── index.html # UI for library system
│ ├── script.js # JS to handle frontend logic
│ ├── style.css # Basic styling
│
└── README.md # This file

---

## 🔧 Prerequisites

- Go 1.18 or higher
- MySQL installed and running
- A database named `library_db`
- A MySQL user with access:  

  Example credentials (you can change them):

username: go_user
password: S3cur3P@ssw0rd

---

## 🛠️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/library-system
cd library-system
```

### 2. Set up the MySQL database
Run the following commands in your MySQL terminal:
```bash 
CREATE DATABASE library_db;
CREATE USER 'go_user'@'localhost' IDENTIFIED BY 'S3cur3P@ssw0rd';
GRANT ALL PRIVILEGES ON library_db.* TO 'go_user'@'localhost';
FLUSH PRIVILEGES;
```

### 3. Run the Go server

go run main.go
### 4. Open the browser
Visit: http://localhost:8080

You can now:

Submit a new book using the form

See all books in the table

⚠️ Disclaimer
This project is created strictly for learning purposes.
It may include hardcoded credentials, lacks validations, and isn't secure for production use.
It will be gradually upgraded as the learning progresses.

### If it is not working

![image](https://i.imgflip.com/4t169s.jpg)
It'll work