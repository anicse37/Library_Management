# 📚 Library Management System

A web-based Library Management System built with **Go (Golang)**, **MySQL**, and **HTML/CSS**. It includes user authentication, role-based access control, and book borrowing features.

---

## 🚀 Features

### 👥 User Roles
- **User**
  - Browse available books
  - Borrow and return books
- **Admin**
  - Add or remove books
  - Remove users
- **Super Admin**
  - Has full control
  - Can approve new admin accounts
  - Register new Admins

### 🔐 Security & Session Management
- **Password hashing** with `bcrypt`
- **Sessions** and **cookies** for persistent login
- Role validation and route protection

---


## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/library-management.git
cd library-management
```
### 2. Set up MySQL

Create a database: library_db
Import tables and constraints (provide .sql file if needed)

### 3. Configure Environment (Optional)
If needed, use a .env file or configure database credentials in the code (variables.go).

### 4. Run the server
```bash
go run main.go
```
 - The server should start on localhost:5050

## 🧪 Development Notes
 - Passwords are securely hashed with bcrypt
 - Users are authenticated via session cookies
 - Only superadmins can approve admin accounts
 - Templates and CSS are organized per page

## 📌 Future Improvements
 - Add pagination and search filters
 - Implement Docker support
 - Add input validation and CSRF protection
 - Improve frontend with JS enhancements
 - Include test coverage using _test.go


📄 License
This project is open-source and free to use.

