# Task Management API (Golang + Gin + PostgreSQL + JWT)

Project ini adalah REST API sederhana untuk manajemen task dengan sistem autentikasi JWT dan relasi user → task.  
Dibangun sebagai portfolio backend untuk melamar pekerjaan Software Engineer di Jepang 🇯🇵.

---

## 🚀 Tech Stack
| Layer | Tools |
|------|-------|
| Backend | Go (Golang) |
| Framework | Gin Web Framework |
| Database ORM | GORM |
| Database | PostgreSQL |
| Auth | JWT (JSON Web Token) |
| Password Security | bcrypt hashing |

---

## 📂 Project Structure
task-api/
│
├── cmd/
│ └── main.go
├── internal/
│ ├── database/
│ │ └── db.go
│ ├── handlers/
│ │ ├── user_handler.go
│ │ └── task_handler.go
│ ├── middleware/
│ │ └── auth_middleware.go
│ ├── models/
│ │ ├── user.go
│ │ └── task.go
│ └── routes/
│ └── routes.go
└── go.mod
