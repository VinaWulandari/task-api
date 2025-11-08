# Task Management API (Golang + Gin + PostgreSQL + JWT + Swagger)

This is a simple and clean **Task Management API** built using **Golang**, **Gin Framework**, **GORM**, and **PostgreSQL**, with **JWT Authentication** and **Swagger API Documentation**.

This project demonstrates:
- Backend architecture design
- User authentication & authorization
- Protected routes using JWT
- Relational data with Users → Tasks
- Clean folder structure following Go best practices

---

## 🚀 Tech Stack
| Layer | Technology |
|------|------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| ORM | GORM |
| Database | PostgreSQL |
| Auth | JWT |
| Documentation | Swagger UI |
| Deployment Ready | Yes |

---

## 📌 Features
- User Registration & Login
- Login returns JWT token for authentication
- Create / Read / Update / Delete Tasks
- Each Task is owned by a User (Authorization)
- Swagger UI for interactive API testing

---

## 🗂️ Project Structure
task-api/
├ cmd/
│ └ main.go
├ internal/
│ ├ handlers/ → HTTP Handlers
│ ├ routes/ → Route definitions
│ ├ services/ → (Optional extension)
│ ├ models/ → Database Models
│ └ database/ → PostgreSQL Connection
├ docs/ → Swagger generated docs
├ .env
└ go.mod

## 🔐 Authentication
All protected API endpoints require Bearer Token:
Authorization: Bearer <your_token>

## 📚 API Documentation (Swagger)
Start server:
```bash
go run cmd/main.go
Open browser: http://localhost:8080/swagger/index.html

🧪 Example Endpoints
Auth
------------------------------------------
Method	Endpoint	Description
POST	/api/register	Create new user
POST	/api/login	Login & receive JWT token

Tasks (Requires Token)
------------------------------------------
Method	Endpoint	Description
GET	/api/tasks	Get user tasks
POST	/api/tasks	Create new task
PUT	/api/tasks/{id}	Update task
DELETE	/api/tasks/{id}	Delete task

----------------------------------------------------
🧑‍💻 Author

Name: Vina Wulandari
Role: QA → Backend Developer (Career Transition)
Location: Indonesia
Goal: Backend Engineer

I am currently improving my backend engineering skills and preparing for professional opportunities.