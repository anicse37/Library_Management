# 📚 Library Management System (Go + MySQL)

This is a basic **Library Management System** built using:

- Go (Golang) for backend
- MySQL for database
- HTML/CSS for frontend
- Gorilla sessions for session management

> ⚠️ This project is created **for learning purposes only**. It is not production-ready.

---

## ✨ Features

- User registration & login
- Role-based access (User, Admin, Super Admin)
- Book listing, searching, and borrowing
- Admin approval for users
- Session-based authentication

---

## 🚀 Getting Started

1. **Install:**
   - Go (v1.18 or later)
   - MySQL

2. **Update DB credentials:**
   Edit the `dns` constant in `main.go` with your own MySQL connection details.

3. **Run the server:**
   ```bash
   go run main.go
Access the app:
Open http://localhost:5050/login in your browser.

🛑 Disclaimer
This code is strictly for learning and demonstration purposes.
It lacks production-level features such as input validation, security hardening, and error handling.
### If it is not working

![image](https://i.imgflip.com/4t169s.jpg)
It'll work